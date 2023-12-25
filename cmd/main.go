package main

import (
	"log"
	"os"
	"os/signal"

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

	log.Println("Adding commands...")
	commands.SlashCommands(s)
	components.ComponentsHandler(s)
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Removing commands...")
	commands.RemoveSlashCommands(s)

	// ** DELETE ALL COMMANDS **
	// applications, err := s.ApplicationCommands(s.State.User.ID, "")
	// if err != nil {
	// 	fmt.Println("Error getting application commands:", err)
	// 	return
	// }

	// for _, v := range applications {
	// 	err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
	// 	if err != nil {
	// 		log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
	// 	}
	// }

	log.Println("Gracefully shutting down.")
}
