package quests

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	progress = PlayerQuest.QuestProgress
)

func generateQuest0(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

func ButtonQuest0(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		label    string
		content  string
		customID string = "quest0_Button"
	)

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
		label = "Advance"
		content = "Scalling the towering tree, in a distance, you see other lonely trees, none as tall as the one you are in. The shrub land seemed to be getting denser the further away you look"
		PlayerQuest.QuestNumber = 1
		customID = "quest_generate"
	default:
		label = "Error" // this is if a random value appears by a bug or whatever
		content = "Error"
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
