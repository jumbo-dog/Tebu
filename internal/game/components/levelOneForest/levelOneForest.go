package levelOneForest

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/pkg/dialog"

	"github.com/bwmarrin/discordgo"
)

var (
	Sticks, Stones int = 0, 0
	paragraph      string
	DisableSticks  bool
	DisableStone   bool
	disableCamp    bool = true
	maxSticks      string
	maxPebbles     string
	MaxResources   string
	dialogs        = dialog.GetDialog("../../internal/questText/quest_001.json")
)

func LevelOneForest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	lastSave, errSave := save.GetSave(i.User.ID)
	if errSave != nil {
		fmt.Println("Error fetching save data:", errSave)
		return
	}

	gatherResources("gather_wood_button", &Sticks, &DisableSticks, i)
	gatherResources("gather_pebbles", &Stones, &DisableStone, i)
	updateParagraph()
	disableCamp = shouldDisableCamp(lastSave.Resources, Sticks, Stones)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: paragraph + "\nWood: " + strconv.Itoa(Sticks) + " " + maxSticks + "\nStone: " + strconv.Itoa(Stones) + " " + maxPebbles + "\n" + MaxResources,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Gather some sticks",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🪵",
							},
							Disabled: DisableSticks,
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
							Disabled: disableCamp,
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
	if Sticks >= 12 || Stones >= 12 {
		paragraph = dialogs[1].DialogText[0]
	}
	if Sticks >= 20 || Stones >= 20 {
		paragraph = dialogs[2].DialogText[0]
	}

	if Sticks > 5 && maxPebbles == "" {
		maxSticks = dialogs[3].DialogText[0]
	}
	if Stones > 5 && maxSticks == "" {
		maxPebbles = dialogs[3].DialogText[0]
	}

	if Sticks == 20 && Stones < 20 {
		maxSticks = dialogs[4].DialogText[0]
	}
	if Stones == 20 && Sticks < 20 {
		maxPebbles = dialogs[4].DialogText[0]
	}

	if Sticks == 20 && Stones == 20 {
		maxSticks = ""
		maxPebbles = ""
		MaxResources = dialogs[5].DialogText[0]
		disableCamp = false
	}
}

func shouldDisableCamp(resources map[string]uint32, Sticks int, Stones int) bool {
	if resources != nil {
		return false
	}
	if resources == nil {
		return true
	}
	return true
}
