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
