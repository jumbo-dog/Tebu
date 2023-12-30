package camp

import (
	"fmt"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/internal/game/components/levelOneForest"

	"github.com/bwmarrin/discordgo"
)

var (
	disableStorage bool
	hasResources   string
)

func GoToCamp(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	resources := levelOneForest.Sticks + levelOneForest.Stones
	disableStorage = true
	if resources > 0 {
		disableStorage = false
	}
	if i.MessageComponentData().CustomID == "store_materials_button" {
		lastSave, errSave := save.GetSave(i.User.ID)
		if errSave != nil {
			if errSave != nil {
				fmt.Println("Error sending direct message:", errSave)
			}
			return
		}
		if lastSave.Resources == nil {
			lastSave.Resources = make(map[string]uint32)
		}
		lastSave.Resources["wood"] += uint32(levelOneForest.Sticks)
		lastSave.Resources["stone"] += uint32(levelOneForest.Stones)
		save.UpdateSave(lastSave)
		levelOneForest.Sticks = 0
		levelOneForest.Stones = 0
		levelOneForest.DisableSticks = false
		levelOneForest.DisableStone = false
		levelOneForest.MaxResources = ""
		disableStorage = true
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Camp's options",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Craft",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🛠",
							},
							CustomID: "craft_button",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Store materials",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "📦",
							},
							Disabled: disableStorage,
							CustomID: "store_materials_button",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Go back to the forest",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🌳",
							},
							CustomID: "goto_forest",
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
