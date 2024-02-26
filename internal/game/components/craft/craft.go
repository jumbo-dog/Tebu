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

	woodenSpear         string = "???"
	woodenSpearId       string = "randomId1"
	woodenSpearDisabled bool
	woodenSpearEmoji    string = "🚫"

	axe         string = "???"
	axeId       string = "randomId2"
	axeDisabled bool
	axeEmoji    string = "🚫"

	pickaxe         string = "???"
	pickaxeId       string = "randomId3"
	pickaxeDisabled bool
	pickaxeEmoji    string = "🚫"
)

func Craft(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	axeDisabled = false
	woodenSpearDisabled = false
	pickaxeDisabled = false
	disableTorch = false
	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error sending direct message:", errSave)
		return
	}
	craftTorch(i, lastSave)
	craftSpear(i, lastSave)
	craftAxe(i, lastSave)
	craftPickaxe(i, lastSave)
	canCraft(lastSave)
	isMaxResources(lastSave)
	save.UpdateSave(lastSave)
	fmt.Println(i.MessageComponentData().CustomID)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Mesa de trabalho:\n " + "Resources: wood: " + strconv.Itoa(int(lastSave.Resources["wood"])) + camp.FullBackpackWood + ", " + "stone: " + strconv.Itoa(int(lastSave.Resources["stone"])) + camp.FullBackpackStone,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Torch (10 Wood)",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🔥",
							},
							Disabled: disableTorch,
							CustomID: "craft_torch",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: woodenSpear,
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: woodenSpearEmoji,
							},
							Disabled: woodenSpearDisabled,
							CustomID: woodenSpearId,
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    axe,
							Disabled: axeDisabled,
							Emoji: discordgo.ComponentEmoji{
								Name: axeEmoji,
							},
							Style:    discordgo.SuccessButton,
							CustomID: axeId,
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    pickaxe,
							Disabled: pickaxeDisabled,
							Emoji: discordgo.ComponentEmoji{
								Name: pickaxeEmoji,
							},
							Style:    discordgo.SuccessButton,
							CustomID: pickaxeId,
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

func craftTorch(i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == "craft_torch" {
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
func craftAxe(i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == "craft_axe" {
		if lastSave.Items == nil {
			lastSave.Items = make(map[string]uint16)
		}
		lastSave.Items["axe"] += 1
		if lastSave.Resources == nil {
			lastSave.Resources = make(map[string]uint32)
		}
		lastSave.Resources["wood"] -= 20
	}
}
func craftPickaxe(i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == "craft_pickaxe" {
		if lastSave.Items == nil {
			lastSave.Items = make(map[string]uint16)
		}
		lastSave.Items["pickaxe"] += 1
		if lastSave.Resources == nil {
			lastSave.Resources = make(map[string]uint32)
		}
		lastSave.Resources["wood"] -= 10
		lastSave.Resources["stone"] -= 20
	}
}
func craftSpear(i *discordgo.InteractionCreate, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == "craft_wooden_spear" {
		if lastSave.Items == nil {
			lastSave.Items = make(map[string]uint16)
		}
		lastSave.Items["wooden_spear"] += 1
		if lastSave.Resources == nil {
			lastSave.Resources = make(map[string]uint32)
		}
		lastSave.Resources["wood"] -= 30
	}
}
func canCraft(lastSave *models.PlayerSave) {
	if lastSave.Resources == nil || lastSave.Resources["wood"] < 10 {
		disableTorch = true
	}
	if lastSave.Resources == nil || lastSave.Resources["wood"] < 30 {
		woodenSpearDisabled = true
	}
	if lastSave.Resources == nil || lastSave.Resources["wood"] < 20 {
		axeDisabled = true
	}
	if lastSave.Resources == nil || lastSave.Resources["wood"] < 10 || lastSave.Resources["stone"] < 20 {
		pickaxeDisabled = true
	}
	if lastSave.Progress.Quest.QuestNumber == 2 {
		axe = "Axe (20 woods)"
		axeEmoji = "🪓"
		axeId = "craft_axe"

		woodenSpear = "Wooden spear (50 woods)"
		woodenSpearEmoji = "⚔️"
		woodenSpearId = "craft_wooden_spear"

		pickaxe = "Pickaxe (10 woods, 20 stones)"
		pickaxeEmoji = "⛏️"
		pickaxeId = "craft_pickaxe"
	}
}
func isMaxResources(lastSave *models.PlayerSave) {
	if lastSave.Resources["wood"] < 50 {
		camp.FullBackpackWood = ""
	}
	if lastSave.Resources["stone"] < 50 {
		camp.FullBackpackStone = ""
	}
}
