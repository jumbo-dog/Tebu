package levelOneForest

import (
	"fmt"
	"strconv"
	"tebu-discord/pkg/dialog"

	"github.com/bwmarrin/discordgo"
)

var (
	Wood, Stones int = 0, 0
	paragraph      string
	DisableWood  bool
	DisableStone   bool
	maxWood        string
	maxStones       string
	MaxResources   string
	dialogs        = dialog.GetDialog("./questText/quest_001.json")
)

func LevelOneForest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	gatherResources("gather_wood_button", &Wood, &DisableWood, i)
	gatherResources("gather_pebbles", &Stones, &DisableStone, i)
	updateParagraph()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: paragraph + "\nWood: " + strconv.Itoa(Wood) + " " + maxWood + "\nStone: " + strconv.Itoa(Stones) + " " + maxStones + "\n" + MaxResources,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Gather some Wood",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🪵",
							},
							Disabled: DisableWood,
							CustomID: "gather_wood_button",
						},
					},
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Gather pebbles",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🪨",
							},
							Disabled: DisableStone,
							CustomID: "gather_pebbles",
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
func gatherResources(customID string, resource *int, disableFlag *bool, i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == customID && *resource < 20 {
		*resource++
	}
	if i.MessageComponentData().CustomID == customID && *resource == 20 {
		*disableFlag = true
	}
}

func updateParagraph() {
	paragraph = dialogs[0].DialogText[0]
	if Wood >= 12 || Stones >= 12 {
		paragraph = dialogs[1].DialogText[0]
	}
	if Wood >= 20 || Stones >= 20 {
		paragraph = dialogs[2].DialogText[0]
	}

	if Wood > 5 && maxStones == "" {
		maxWood = dialogs[3].DialogText[0]
	}
	if Stones > 5 && maxWood == "" {
		maxStones = dialogs[3].DialogText[0]
	}

	if Wood == 20 && Stones < 20 {
		maxWood = dialogs[4].DialogText[0]
	}
	if Stones == 20 && Wood < 20 {
		maxStones = dialogs[4].DialogText[0]
	}

	if Wood == 20 && Stones == 20 {
		maxWood = ""
		maxStones = ""
		MaxResources = dialogs[5].DialogText[0]
	}
}
