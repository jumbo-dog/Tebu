package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	config "tebu-discord/database/config"
	commands "tebu-discord/internal/commands/entity"
	components "tebu-discord/internal/game/components/entity"

	"github.com/bwmarrin/discordgo"
	"github.com/julienschmidt/httprouter"
)

var (
	s            *discordgo.Session
	mainBotToken = "BOT_TOKEN"
	// testBotToken = "BOT_TOKEN_TEST"
)

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Getenv(mainBotToken))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func newRouter() *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/live", statusCheck())
	return mux
}

type Live struct {
	IsLive bool `json:"islive"`
}

func statusCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		live := Live{
			IsLive: true,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(live); err != nil {
			panic(err)
		}
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v\n", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	commands.CreateSlashCommands(s)
	config.ConnectDatabase()
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			commands.HandleSlashCommands(s, i)
		case discordgo.InteractionMessageComponent:
			components.HandleComponents(s, i)
		}
	})

	srv := &http.Server{
		Addr:    ":10101",
		Handler: newRouter(),
	}

	idleConnsClosed := make(chan struct{})

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Printf("fatal http server failed to start: %v", err)
		}
	}

	defer s.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	<-idleConnsClosed
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("http server shutdown error %v", err)
	}

	close(idleConnsClosed)
	commands.RemoveSlashCommands(s)
	config.DisconnectDatabase()

	log.Println("Gracefully shutting down.")
}
