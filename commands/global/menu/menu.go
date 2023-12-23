package direct

import (
	"log"
	"tebu-discord/commands/global/menu/handler"
	helper "tebu-discord/helper/env"
	service "tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

type menu struct {
	session service.SessionService
}

type menuInterface interface {
	StartMenu(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func New(session service.SessionService) menuInterface {
	return &menu{
		session: session,
	}
}

var (
	commands = []discordgo.ApplicationCommand{
		{
			Name: "menu",
			Type: discordgo.UserApplicationCommand,
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"menu": handler.MenuHandler,
	}
	created = false
)

func (m *menu) StartMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandsHandlers["menu"]; ok && created == false {
		h(s, i)
		created = true
	}

	cmdIDs := make(map[string]string, len(commands))
	for _, cmd := range commands {
		rcmd, err := s.ApplicationCommandCreate(helper.GetEnvValue("APP_ID"), "", &cmd)
		if err != nil {
			log.Fatalf("Cannot create slash command %q: %v", cmd.Name, err)
		}
		cmdIDs[rcmd.ID] = rcmd.Name
	}
}
