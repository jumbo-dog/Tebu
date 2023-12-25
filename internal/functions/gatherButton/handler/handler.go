package handler

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	points int
)

func IncrementButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	points++
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "value: " + strconv.Itoa(points),
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label:    "Gather wood",
					Style:    discordgo.SuccessButton,
					CustomID: "button_quest0_01",
				}},
			}},
		},
	})
	if err != nil {
		log.Fatalf("Error creating increment button: %s", err)
	}
}
