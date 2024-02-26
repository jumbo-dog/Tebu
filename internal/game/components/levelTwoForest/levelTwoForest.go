package levelTwoForest

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/internal/game/components/combatsys"
	"tebu-discord/pkg/dialog"

	"github.com/bwmarrin/discordgo"
)

var (
	dialogs        = dialog.GetDialog("./questText/quest_002.json")
	paragraph      = dialogs[0].DialogText[0]
	canChop        bool
	canMine        bool
	forestProgress int64 = 1
	Wood, Stones   int   = 0, 0
	DisableWood    bool
	DisableStone   bool
	maxWood        string
	maxStones      string
	MaxResources   string
	disableExplore bool
	buttonId       string = "explore"
)

func LevelTwoForest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {

	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error sending direct message:", errSave)
		return
	}
	canChop = false
	canMine = false
	hasTools(i, lastSave)
	combatsys.GetCombatStats(combatsys.FOREST_2, i.User.ID, "explore", "goto_camp")
	forestProgression(i)
	gatherMaterials(i, "chop_logs", &canChop, &Wood, lastSave)
	gatherMaterials(i, "mine_rocks", &canMine, &Stones, lastSave)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: paragraph + "\nWood: " + strconv.Itoa(Wood) + " " + maxWood + "\nStone: " + strconv.Itoa(Stones) + " " + maxStones + "\n" + MaxResources,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Explore",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🧭",
							},
							CustomID: buttonId,
							Disabled: disableExplore,
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Chop logs",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🪓",
							},
							Disabled: canChop,
							CustomID: "chop_logs",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Mine rocks",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "⛏️",
							},
							Disabled: canMine,
							CustomID: "mine_rocks",
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

func hasTools(i *discordgo.InteractionCreate, save *models.PlayerSave) {
	if save.Items["axe"] == 0 {
		canChop = true
	}
	if save.Items["pickaxe"] == 0 {
		canMine = true
	}
}

func forestProgression(i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "explore" {
		forestProgress += 1
		if forestProgress == 2 {
			paragraph = dialogs[1].DialogText[0]
		}
		if forestProgress == 3 {
			paragraph = dialogs[2].DialogText[0]
			buttonId = dialogs[2].ButtonLabel[0]
		}
		if forestProgress == 4 {
			paragraph = dialogs[3].DialogText[0]
			buttonId = dialogs[3].ButtonLabel[0]
		}
		if forestProgress == 5 {
			paragraph = dialogs[4].DialogText[0]
		}
		if forestProgress == 6 {
			paragraph = dialogs[5].DialogText[0]
		}
	}

}

func gatherMaterials(i *discordgo.InteractionCreate, customID string, disableFlag *bool, resource *int, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == customID {
		*resource += 3
	}
	if *resource > 20 && lastSave.Items["backpack"] < 1 {
		*disableFlag = true
		*resource = 20
	}
}
