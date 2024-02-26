package entity

import (
	"tebu-discord/internal/game/components/camp"
	"tebu-discord/internal/game/components/chooseForest"
	"tebu-discord/internal/game/components/combatsys"
	"tebu-discord/internal/game/components/craft"
	"tebu-discord/internal/game/components/dungeon"
	"tebu-discord/internal/game/components/levelOneForest"
	"tebu-discord/internal/game/components/levelTwoForest"
	"tebu-discord/internal/game/quests"
	entity "tebu-discord/internal/game/quests/entity"

	"github.com/bwmarrin/discordgo"
)

var (
	componentsHandlers = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		"gather_wood_button":     levelOneForest.LevelOneForest,
		"gather_pebbles":         levelOneForest.LevelOneForest,
		"chop_logs":              levelTwoForest.LevelTwoForest,
		"mine_rocks":             levelTwoForest.LevelTwoForest,
		"goto_forest":            levelOneForest.LevelOneForest,
		"goto_forest2":           levelTwoForest.LevelTwoForest,
		"explore":                levelTwoForest.LevelTwoForest,
		"init_attack":            combatsys.HandleCombat,
		"attack_button":          combatsys.HandleAttacks,
		"goto_camp":              camp.GoToCamp,
		"store_materials_button": camp.GoToCamp,
		"goto_craftbench":        craft.Craft,
		"craft_torch":            craft.Craft,
		"craft_axe":              craft.Craft,
		"craft_pickaxe":          craft.Craft,
		"craft_wooden_spear":     craft.Craft,
		"explore_dungeon":        dungeon.StartDungeon,
		"choose_forest":          chooseForest.ChooseWhereToGo,
		"quest_generate":         entity.GenerateQuest,
		"quest0_Button":          quests.ButtonQuest0,
	}
)

func HandleComponents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
		h(s, i)
	}
}
