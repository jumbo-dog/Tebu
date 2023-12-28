package quests

import (
	"log"
	"tebu-discord/database/models"

	"github.com/bwmarrin/discordgo"
)

func GenerateQuest1(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "You suddenly feel a primal urge to start a civilization, the next step should be obvious...",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label: "Gather some sticks",
					Style: discordgo.SuccessButton,
					Emoji: discordgo.ComponentEmoji{
						Name: "🪵",
					},
					CustomID: "gather_wood_button",
				}},
			}},
		},
	})
	if err != nil {
		log.Fatalf("Error creating button: %v", err)
	}
}
