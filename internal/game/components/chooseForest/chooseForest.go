package chooseForest

import (
	"fmt"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"

	"github.com/bwmarrin/discordgo"
)

func ChooseWhereToGo(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	save.UpdateSave(&models.PlayerSave{
		DiscordId: i.User.ID,
		Progress: &models.Progress{
			Quest: &models.Quest{
				QuestNumber:   2,
				QuestProgress: 0,
			},
		},
	})
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Choose where you want to go",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Forest",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🌳",
							},
							CustomID: "goto_forest",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Go deeper into the forest using the torch",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🧛",
							},
							CustomID: "goto_forest2",
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("Error creating increment button: %v \n", err)
	}
}
