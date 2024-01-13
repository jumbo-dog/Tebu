package main

import (
	"log"
	"os"
	"os/signal"

	config "tebu-discord/database/config"
	commands "tebu-discord/internal/commands/entity"
	components "tebu-discord/internal/game/components/entity"

	"github.com/bwmarrin/discordgo"
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

	defer s.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	commands.RemoveSlashCommands(s)
	config.DisconnectDatabase()

	log.Println("Gracefully shutting down.")
}
