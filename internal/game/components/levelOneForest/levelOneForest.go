package levelOneForest

import (
	"fmt"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"

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
	paragraph = "The trees surround you gently, you can feel the breeze passing through the forest"
	if Sticks >= 20 || Stones >= 20 {
		paragraph = "The darkened forest unsettles you, prompting a decision to turn back. The unease lingers as you retreat, contemplating a return with a torch"
	}
	if Sticks >= 12 || Stones >= 12 {
		paragraph = "The forest envelops you, a musky scent in the air. Strings of light filter through dense foliage. Each step feels like a venture into a realm alive with ancient secrets."
	}

	if Sticks > 5 && maxPebbles == "" {
		maxSticks = "*(Your hands are getting tired, maybe there is a better way to do this)*"
	}
	if Stones > 5 && maxSticks == "" {
		maxPebbles = "*(Your hands are getting tired, maybe there is a better way to do this)*"
	}

	if Sticks == 20 && Stones < 20 {
		maxSticks = "*(Your arms are getting heavy)*"
	}
	if Stones == 20 && Sticks < 20 {
		maxPebbles = "*(Your arms are getting heavy)*"
	}

	if Sticks == 20 && Stones == 20 {
		maxSticks = ""
		maxPebbles = ""
		MaxResources = "*(You feel like one more flower would make you collapse)*"
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
