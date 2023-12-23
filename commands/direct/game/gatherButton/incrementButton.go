package game

import (
	"log"
	"tebu-discord/commands/direct/game/gatherButton/handler"
	"tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

type gather struct {
	session service.SessionService
}

type gatherInterface interface {
	GatherButton(s *discordgo.Session, m *discordgo.MessageCreate)
}

func New(session service.SessionService) gatherInterface {
	return &gather{
		session: session,
	}
}

var (
	button = discordgo.Button{
		Label:    "Gather wood",
		Style:    discordgo.SuccessButton,
		CustomID: "button_quest0_01",
	}
)

func (g *gather) GatherButton(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "world")
	}

	if m.Content == "plays" {
		incrementButton(s, m)
	}
}

func incrementButton(s *discordgo.Session, m *discordgo.MessageCreate) {
	userChannel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		s.ChannelMessageSend(userChannel.ID, "Error displaying message")
		log.Fatalf("Error creating channel: %s", err)
	}
	s.ChannelMessageSendComplex(userChannel.ID, &discordgo.MessageSend{
		Content: "Click the button below!",
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{button},
			}},
	})
	s.AddHandler(handler.IncrementButtonHandler)
}
