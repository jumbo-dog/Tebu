package game

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func IncrementButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			CustomID: "increment_button",
			Title:    "Click Me!",
			Flags:    discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Button",
							Style:    discordgo.SuccessButton,
							CustomID: "button_quest0_01",
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("Error creating increment button: %s", err)
	}
}
