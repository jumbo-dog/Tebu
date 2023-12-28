package entity

import (
	"tebu-discord/database/models"
	gatherwood "tebu-discord/internal/game/components/gatherWood"
	"tebu-discord/internal/game/quests"
	generateQuest "tebu-discord/internal/game/quests/entity"
	questEntity "tebu-discord/internal/game/quests/entity"

	"github.com/bwmarrin/discordgo"
)

var (
	playersave = questEntity.PlayerSave

	componentsHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate, ...*models.PlayerSave){
		// Components
		"gather_wood_button": gatherwood.GatherWoodButton,

		// Menu components
		"quest_generate": generateQuest.GenerateQuest,

		// Quests components
		"quest0_Button": quests.ButtonQuest0,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i, &playersave)
	}
}
