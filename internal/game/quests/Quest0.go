package quests

import (
	"fmt"
	"log"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/pkg/dialog"
	"tebu-discord/pkg/timer"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	progress uint8
	dialogs  = dialog.GetDialog("./questText/quest_000.json")
)

func GenerateQuest0(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "Game started!\n **Remember: read the prompt and the button text**",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})

	paragraph := dialogs[0].DialogText[0]
	time.Sleep(time.Second * 1) // Value I found to be long enough to start quest but on to long
	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Flags:   discordgo.MessageFlagsEphemeral,
		Content: paragraph,
	})
	if err != nil {
		log.Fatalf("Error creating follow up messsage: %v", err)
	}

	time.Sleep(timer.GenerateQuestTime(paragraph))

	paragraph = dialogs[1].DialogText[0]

	s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
		Content: &paragraph,
	})
	time.Sleep(timer.GenerateQuestTime(paragraph))

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
		fmt.Printf("Unable to get save %s\n", err)
	}
	progress++ // We need to sync this with the db
	cloudSave.Progress.Quest.QuestProgress = progress
	switch progress {
	case 1:
		// The 0ith and 1th is used in the function above
		label = dialogs[2].ButtonLabel[0]
		content = dialogs[2].DialogText[0]
	case 2:
		label = dialogs[3].ButtonLabel[0]
		content = dialogs[3].DialogText[0]
	case 3:
		label = dialogs[4].ButtonLabel[0]
		content = dialogs[4].DialogText[0]
	case 4:
		label = dialogs[5].ButtonLabel[0]
		content = dialogs[5].DialogText[0]
	case 5:
		label = dialogs[6].ButtonLabel[0]
		content = dialogs[6].DialogText[0]

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
