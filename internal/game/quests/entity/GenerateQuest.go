package entity

import (
	"log"
	"tebu-discord/database/controller/save"
	"tebu-discord/database/models"
	"tebu-discord/internal/game/quests"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	PlayerSave     models.PlayerSave
	questsHandlers = map[uint16]func(*discordgo.Session, *discordgo.InteractionCreate, ...*models.PlayerSave){
		0: quests.GenerateQuest0,
		1: quests.GenerateQuest1,
	}
)

func GenerateQuest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	var questNumber uint16

	PlayerSave, err := save.GetSave(i.User.ID)
	if err == mongo.ErrNoDocuments {
		PlayerSave.DiscordId = i.User.ID
		PlayerSave.LastUsername = i.User.Username
		PlayerSave.Progress.Quest.QuestNumber = 0
		save.CreateSave(PlayerSave)
	}

	if err != nil && err != mongo.ErrNoDocuments {
		log.Fatalf("Error generating quest: %v", err)
	}
	questNumber = PlayerSave.Progress.Quest.QuestNumber
	if h, ok := questsHandlers[questNumber]; ok {
		h(s, i, PlayerSave)
	}
}
