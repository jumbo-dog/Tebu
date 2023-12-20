package direct

import (
	service "tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

type DirectMenu struct {
	session service.SessionService
}

// NOT WORKING YET
func (d *DirectMenu) startMenu() {
	d.session.StartSession().AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world")
		}
	})
}
