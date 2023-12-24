package game

import (
	"log"
	"tebu-discord/internal/functions/gatherButton/handler"

	"github.com/bwmarrin/discordgo"
)

var (
	button = discordgo.Button{
		Label:    "Gather wood",
		Style:    discordgo.SuccessButton,
		CustomID: "button_quest0_01",
	}
)

func IncrementButton(s *discordgo.Session, m *discordgo.MessageCreate) {
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
