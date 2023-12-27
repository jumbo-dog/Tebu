package entity

import (
	gatherwood "tebu-discord/internal/game/components/gatherWood"
	"tebu-discord/internal/game/quests"

	"github.com/bwmarrin/discordgo"
)

var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Components
		"gather_wood_button": gatherwood.GatherWoodButton,

		// Quests
		"quest_generate": quests.GenerateQuest,
		"quest0_Button":  quests.ButtonQuest0,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i)
	}
}
