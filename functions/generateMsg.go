package functions

import (
	game "tebu-discord/commands/direct/game"
	service "tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

type DirectMenu struct {
	session service.SessionService
}

func MessageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "world")
	}

	if m.Content == "plays" {
		game.IncrementButton(s, m)
	}
}
