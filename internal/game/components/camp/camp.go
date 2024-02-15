package camp

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/internal/game/components/levelOneForest"

	"github.com/bwmarrin/discordgo"
)

var (
	disableStorage    bool = true
	FullBackpackWood       = ""
	FullBackpackStone      = ""
	gotoWhere              = "goto_forest"
)

func GoToCamp(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {

	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error sending direct message:", errSave)
		return
	}
	checkResourses(lastSave)
	storeMaterials("store_materials_button", i, lastSave)
	if lastSave.Items["torch"] != 0 {
		gotoWhere = "choose_forest"
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Camp's options:\nResources: wood: " + strconv.Itoa(int(lastSave.Resources["wood"])) + FullBackpackWood + ", " + "stone: " + strconv.Itoa(int(lastSave.Resources["stone"])) + FullBackpackStone,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Craft",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🛠️",
							},
							CustomID: "goto_craftbench",
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
							CustomID: gotoWhere,
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

func checkResourses(lastSave *models.PlayerSave) {
	FullBackpackWood = ""
	FullBackpackStone = ""
	if lastSave.Resources["wood"] == 50 {
		FullBackpackWood = " *(MAX)*"
	}
	if lastSave.Resources["stone"] == 50 {
		FullBackpackStone = " *(MAX)*"
	}
	resources := levelOneForest.Wood + levelOneForest.Stones
	if resources > 0 {
		disableStorage = false
	}
}

func isBiggerThanBackpack(lastSave *models.PlayerSave, SavedWood uint32, SavedStone uint32) {
	if lastSave.Resources == nil {
		lastSave.Resources = make(map[string]uint32)
	}
	if lastSave.Resources != nil && int(SavedWood)+levelOneForest.Wood < 50 {
		lastSave.Resources["wood"] += uint32(levelOneForest.Wood)
	} else {
		lastSave.Resources["wood"] = 50
		FullBackpackWood = " *(MAX)*"
	}
	if lastSave.Resources != nil && int(SavedStone)+levelOneForest.Stones < 50 {
		lastSave.Resources["stone"] += uint32(levelOneForest.Wood)
	} else {
		lastSave.Resources["stone"] = 50
		FullBackpackStone = " *(MAX)*"
	}
}

func storeMaterials(customID string, i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	SavedWood := lastSave.Resources["wood"]
	SavedStone := lastSave.Resources["stone"]
	if i.MessageComponentData().CustomID == customID {
		isBiggerThanBackpack(lastSave, SavedWood, SavedStone)
		save.UpdateSave(lastSave)
		resetForest()
	}
}

func resetForest() {
	levelOneForest.Wood = 0
	levelOneForest.Stones = 0
	levelOneForest.DisableWood = false
	levelOneForest.DisableStone = false
	levelOneForest.MaxResources = ""
	disableStorage = true
}
