package handler

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	userChannel *discordgo.Channel
	points      int
	button      = discordgo.Button{
		Label:    "Gather wood",
		Style:    discordgo.SuccessButton,
		CustomID: "button_quest0_01",
	}
)

func IncrementButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	points++
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Click the button below:\nWood: " + strconv.Itoa(points),
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{button},
			}},
		},
	})
	if err != nil {
		log.Fatalf("Error creating increment button: %s", err)
		s.ChannelMessageSend(userChannel.ID, "Error displaying message")
	}
}
