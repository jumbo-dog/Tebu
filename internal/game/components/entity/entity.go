package entity

import (
	"tebu-discord/internal/game/components/camp"
	"tebu-discord/internal/game/components/chooseForest"
	"tebu-discord/internal/game/components/craft"
	"tebu-discord/internal/game/components/levelOneForest"
	"tebu-discord/internal/game/components/levelTwoForest"
	"tebu-discord/internal/game/quests"
	entity "tebu-discord/internal/game/quests/entity"

	"github.com/bwmarrin/discordgo"
)

var (
	componentsHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		// Components
		"gather_wood_button": levelOneForest.LevelOneForest,
		"gather_pebbles":     levelOneForest.LevelOneForest,
		"chop_logs":          levelTwoForest.LevelTwoForest,
		"mine_rocks":         levelTwoForest.LevelTwoForest,

		"goto_forest":     levelOneForest.LevelOneForest,
		"goto_forest2":    levelTwoForest.LevelTwoForest,
		"goto_camp":       camp.GoToCamp,
		"goto_craftbench": craft.Craft,
		"explore":         levelTwoForest.LevelTwoForest,

		"store_materials_button": camp.GoToCamp,
		"create_torch":           craft.Craft,
		"choose_forest":          chooseForest.ChooseWhereToGo,
		// Menu components
		"quest_generate": entity.GenerateQuest,

		// Quests components
		"quest0_Button": quests.ButtonQuest0,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i)
	}
}
