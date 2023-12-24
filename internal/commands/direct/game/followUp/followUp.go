package followup

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func FollowUp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Followup messages are basically regular messages (you can create as many of them as you wish)
	// but work as they are created by webhooks and their functionality
	// is for handling additional messages after sending a response.

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			// Note: this isn't documented, but you can use that if you want to.
			// This flag just allows you to create messages visible only for the caller of the command
			// (user who triggered the command)
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: "Surprise!",
		},
	})
	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "Followup message has been created, after 5 seconds it will be edited",
	})
	if err != nil {
		s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})
		return
	}
	time.Sleep(time.Second * 5)

	content := "Now the original message is gone and after 10 seconds this message will ~~self-destruct~~ be deleted."
	s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
		Content: &content,
	})

	time.Sleep(time.Second * 10)

	s.FollowupMessageDelete(i.Interaction, msg.ID)

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "For those, who didn't skip anything and followed tutorial along fairly, " +
			"take a unicorn :unicorn: as reward!\n" +
			"Also, as bonus... look at the original interaction response :D",
	})
}
