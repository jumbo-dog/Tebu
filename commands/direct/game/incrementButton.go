package game

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	userChannel *discordgo.Channel
	err         error
	points      int

	button = discordgo.Button{
		Label:    "Gather wood",
		Style:    discordgo.SuccessButton,
		CustomID: "button_quest0_01",
	}
)

func IncrementButton(s *discordgo.Session, m *discordgo.MessageCreate) {
	userChannel, err = s.UserChannelCreate(m.Author.ID)

	if err != nil {
		s.ChannelMessageSend(userChannel.ID, "Error displaying message")
		log.Println(err)
	}
	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}
	s.ChannelMessageSendComplex(userChannel.ID, &discordgo.MessageSend{
		Content:    "Click the button below!",
		Components: []discordgo.MessageComponent{&actionRow},
	})
	s.AddHandler(incrementButtonHandler)
}

func incrementButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	val := "Click the button below:\nWood: " + strconv.Itoa(points)
	points++

	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content:    val,
			Components: []discordgo.MessageComponent{actionRow},
		},
	})
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(userChannel.ID, "Error displaying message")
	}
}
