package entity

import (
	"fmt"
	"log"
	menu "tebu-discord/internal/commands/global/menu"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "menu",
			Description: "Menu with the main options of the bot",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"menu": menu.StartMenu,
	}
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
)

func CreateSlashCommands(s *discordgo.Session) {
	log.Println("Adding commands...")
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

func HandleSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

// NOT USED
// DELETES ALL COMMANDS FROM A USER, ONLY PRIVATE COMMANDS
// To delete guild commands insert guild id
func DeleteAllCommands(s *discordgo.Session) {
	applications, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		fmt.Println("Error getting application commands:", err)
		return
	}

	for _, v := range applications {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
