package quests

import (
	"log"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	progress uint8
)

func GenerateQuest0(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Game started!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	time.Sleep(time.Second * 1) // Value I found to be long enough to start quest but on to long
	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Flags:   discordgo.MessageFlagsEphemeral,
		Content: "All of a sudden, a bright blue light comes shinning at your eyes, illuminating a canvas of clouds that dance in the skies to the most beautiful melody. The air hums in nature's sweet symphony, you catch the invigorating scent of fresh grass.",
	})
	if err != nil {
		log.Fatalf("Error creating follow up messsage: %v", err)
	}

	time.Sleep(time.Second * 1)

	paragraph2 := "All of a sudden, a bright blue light comes shinning at your eyes, illuminating a canvas of clouds that dance in the skies to the most beautiful melody. The air hums in nature's sweet symphony, you catch the invigorating scent of fresh grass.\nGetting up, the landscape unfolds before you, a vast expanse of lush green plains stretching beyond the horizon. The world flares out like a blank sheet of paper, inviting you to explore its mysteries and secrets."

	s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
		Content: &paragraph2,
	})

	time.Sleep(time.Second * 1)

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Flags: discordgo.MessageFlagsEphemeral,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Explore",
						CustomID: "quest0_Button",
					},
				},
			},
		},
	})
}

func ButtonQuest0(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	var (
		label    string
		content  string
		customID string = "quest0_Button"
	)

	cloudSave, err := save.GetSave(i.User.ID)
	if err != nil {

	}
	progress++
	switch progress {
	case 1:
		label = "Explore"
		content = "You run towards the sun, but all you find is an endless sea of grass"
	case 2:
		label = "Explore"
		content = "Nothing new seems to appear at your sight"
	case 3:
		label = "Explore"
		content = "Wind rattles and shakes the vegetation around you, very far something different catches your attention"
	case 4:
		label = "Climb tree"
		content = "A majestic tall tree, rooted to the ever so green grass, you notice rocks shattered within it's roots, while fungi grows within the gap of the base of the tree and the soil."
	case 5:
		label = "Climb down"
		content = "Scalling the towering tree, in a distance, you see other lonely trees, none as tall as the one you are in. The shrub land seemed to be getting denser the further away you look"

		cloudSave.Progress.Quest.QuestNumber = 1
		save.UpdateSave(cloudSave)

		customID = "quest_generate"
	default:
		label = "Error"
		content = "Error"
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label:    label,
					Style:    discordgo.PrimaryButton,
					CustomID: customID,
				}},
			}},
		},
	})
	if err != nil {
		log.Fatalf("Error creating increment button: %s", err)
	}
}
