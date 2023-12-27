package quests

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func generateQuest1(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Release your primal instincts and break the roots of the tree",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label:    "Break the roots of the tree",
					Style:    discordgo.SuccessButton,
					CustomID: "gather_wood_button",
				}},
			}},
		},
	})
	if err != nil {
		log.Fatalf("Error creating button: %v", err)
	}
}
