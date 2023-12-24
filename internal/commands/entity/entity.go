package entity

import (
	"log"
	followup "tebu-discord/internal/commands/direct/game/followUp"
	basiccomandfiles "tebu-discord/internal/commands/global/basic-command-files"
	menu "tebu-discord/internal/commands/global/menu"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command-with-files",
			Description: "Basic command with files",
		},
		{
			Name:        "followups",
			Description: "Followup messages",
		},
		{
			Name:        "menu",
			Description: "Menu with the main options of the bot",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"followups":                followup.FollowUp,
		"basic-command-with-files": basiccomandfiles.BasicComandsFile,
		"menu":                     menu.StartMenu,
	}
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
	created            = false
)

func SlashCommands(s *discordgo.Session) {
	log.Println("Adding commands...")
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok && created == false {
			h(s, i)
			created = true
		}
	})
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func RemoveSlashCommands(s *discordgo.Session) {
	log.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

}
