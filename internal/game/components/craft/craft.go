package craft

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/internal/game/components/camp"

	"github.com/bwmarrin/discordgo"
)

var (
	disableTorch bool
)

func Craft(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	disableTorch = false
	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error sending direct message:", errSave)
		return
	}
	buyTorch(i, lastSave)
	canBuy(lastSave)
	isMaxResources(lastSave)
	save.UpdateSave(lastSave)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Mesa de trabalho:\n " + "Resources: wood: " + strconv.Itoa(int(lastSave.Resources["wood"])) + camp.FullBackpackWood + ", " + "stone: " + strconv.Itoa(int(lastSave.Resources["stone"])) + camp.FullBackpackStone,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Torch",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🔥",
							},
							Disabled: disableTorch,
							CustomID: "create_torch",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "???",
							Style:    discordgo.SuccessButton,
							Disabled: true,
							CustomID: "whatthefrick",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "???",
							Disabled: true,
							Style:    discordgo.SuccessButton,
							CustomID: "OOOOOOOOOOOOOOOOOOOOHMYGAH",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Go to camp",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "⛺",
							},
							CustomID: "goto_camp",
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

func buyTorch(i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == "create_torch" {
		if lastSave.Items == nil {
			lastSave.Items = make(map[string]uint16)
		}
		lastSave.Items["torch"] += 1
		if lastSave.Resources == nil {
			lastSave.Resources = make(map[string]uint32)
		}
		lastSave.Resources["wood"] -= 10
	}
}
func canBuy(lastSave *models.PlayerSave) {
	if lastSave.Resources == nil || lastSave.Resources["wood"] < 10 {
		disableTorch = true
	}
}
func isMaxResources(lastSave *models.PlayerSave) {
	if lastSave.Resources["wood"] < 100 {
		camp.FullBackpackWood = ""
	}
	if lastSave.Resources["stone"] < 100 {
		camp.FullBackpackStone = ""
	}
}
