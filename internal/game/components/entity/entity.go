package entity

import (
	"tebu-discord/database/models"
	"tebu-discord/internal/game/components/camp"
	"tebu-discord/internal/game/components/craft"
	levelOneForest "tebu-discord/internal/game/components/levelOneForest"
	"tebu-discord/internal/game/quests"
	entity "tebu-discord/internal/game/quests/entity"

	"github.com/bwmarrin/discordgo"
)

var (
	playersave = entity.PlayerSave

	componentsHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate, ...*models.PlayerSave){
		// Components
		"gather_wood_button":     levelOneForest.LevelOneForest,
		"gather_pebbles":         levelOneForest.LevelOneForest,
		"goto_forest":            levelOneForest.LevelOneForest,
		"goto_camp":              camp.GoToCamp,
		"store_materials_button": camp.GoToCamp,
		"goto_craftbench":        craft.Craft,
		"create_torch":           craft.Craft,
		// Menu components
		"quest_generate": entity.GenerateQuest,

		// Quests components
		"quest0_Button": quests.ButtonQuest0,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i, &playersave)
	}
}
