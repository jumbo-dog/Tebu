package entity

import (
	"tebu-discord/internal/functions/gatherButton/handler"

	"github.com/bwmarrin/discordgo"
)

var (
	componentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"button_quest0_01": handler.IncrementButtonHandler,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i)
	}
}
