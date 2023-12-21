package direct

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	chnID             string
	incrementedNumber int
)

func IncrementButton(s *discordgo.Session, channelID string) {
	button := discordgo.Button{
		Label:    "Click me!",
		Style:    discordgo.SuccessButton,
		CustomID: "exampleButton",
	}

	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}
	s.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content:    "Click the button below!",
		Components: []discordgo.MessageComponent{&actionRow},
	})
	chnID = channelID
	s.AddHandler(ReactionAdd)
}

func ReactionAdd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	val := strconv.Itoa(incrementedNumber)
	incrementedNumber += 1

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Your val: " + val,
		},
	})
	if err != nil {
		log.Println(err)
	}
}
