package levelTwoForest

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"

	"github.com/bwmarrin/discordgo"
)

var (
	paragraph      = "As you step into the dimly lit depths of the forest, a sense of unease washes over you. The eerie silence is occasionally interrupted by the faint rustling of leaves and the crackling of twigs underfoot."
	hasAxe         bool
	hasPick        bool
	forestProgress int64
	Wood, Stones   int = 0, 0
	DisableWood    bool
	DisableStone   bool
	maxWood        string
	maxStones      string
	MaxResources   string
	disableExplore bool
)

func hasTools(i *discordgo.InteractionCreate, save *models.PlayerSave) {
	if save.Items["axe"] > 0 {
		hasAxe = true
	}
	if save.Items["pickaxe"] > 0 {
		hasPick = true
	}
}

func forestProgression(i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "explore" {
		forestProgress += 1
		if forestProgress >= 10 {
			forestProgress = 10
			disableExplore = true
		}
		if forestProgress == 2 {
			paragraph = "With each step, the undergrowth thickens, obscuring your path and heightening your apprehension."
		}
		if forestProgress == 3 {
			paragraph = "A sudden rustle in the bushes ahead shatters the tranquility, sending a shiver down your spine."
		}
		if forestProgress == 4 {
			paragraph = "Pressing onward, you delve deeper into the wilderness, the encroaching shadows a constant reminder of the dangers lurking in the darkness."
		}
		if forestProgress == 5 {
			paragraph = "And then, amidst the twisted roots and ancient trees, you stumble upon an ancient archway—a gateway to realms unknown."
		}
		if forestProgress == 6 {
			paragraph = "With a mixture of trepidation and curiosity, you cross the threshold, knowing that within the depths of the dungeon await both mysteries untold and perils unimagined."
		}
	}

}

func gatherMaterials(i *discordgo.InteractionCreate, customID string, resource *int, lastSave *models.PlayerSave) {
	if i.MessageComponentData().CustomID == customID {
		*resource += 3
	}
	if *resource > 20 && lastSave.Items["backpack"] < 1 {
		*resource = 20
	}
}

func LevelTwoForest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {

	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error sending direct message:", errSave)
		return
	}
	hasTools(i, lastSave)
	forestProgression(i)
	gatherMaterials(i, "chop_logs", &Wood, lastSave)
	gatherMaterials(i, "mine_rocks", &Stones, lastSave)
	fmt.Println(forestProgress)

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
							CustomID: "explore",
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
							Disabled: hasAxe,
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
							Disabled: hasPick,
							CustomID: "mine_rocks",
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
