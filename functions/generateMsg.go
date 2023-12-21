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
		s.ChannelMessage(m.ChannelID, "world")
	}

	if m.Content == "button" {
		game.IncrementButton(s, m.ChannelID)
	}
}
