package quests

import (
	"tebu-discord/database/models"

	"github.com/bwmarrin/discordgo"
)

var (
	PlayerQuest models.Quest // NOTE: placeholder value, needs to get value from db

	questsHandlers = map[uint16]func(*discordgo.Session, *discordgo.InteractionCreate){
		0: generateQuest0,
		1: generateQuest1,
	}
)

func GenerateQuest(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := questsHandlers[PlayerQuest.QuestNumber]; ok {
		h(s, i)
	}
}
