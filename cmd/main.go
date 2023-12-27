package main

import (
	"log"
	"os"
	"os/signal"

	config "tebu-discord/database/config"
	saveInformation "tebu-discord/database/controller/save"
	commands "tebu-discord/internal/commands/entity"
	components "tebu-discord/internal/components/entity"
	helper "tebu-discord/internal/helper/env"

	"github.com/bwmarrin/discordgo"
)

var (
	s            *discordgo.Session
	mainBotToken = "BOT_TOKEN"
	testBotToken = "BOT_TOKEN_TEST"
)

func init() {
	var err error
	s, err = discordgo.New("Bot " + helper.GetEnvValue(mainBotToken, "../../.env"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	commands.CreateSlashCommands(s)
	config.ConnectDatabase()
	saveInformation.DeleteSave(235)
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
