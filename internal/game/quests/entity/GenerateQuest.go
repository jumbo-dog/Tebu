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
	questsHandlers = map[uint16]func(*discordgo.Session, *discordgo.InteractionCreate){
		0: quests.GenerateQuest0,
		1: quests.GenerateQuest1,
	}
)

func GenerateQuest(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	var questNumber uint16

	PlayerSave, err := save.GetSave(i.User.ID)
	if err == mongo.ErrNoDocuments {
		PlayerSave.LastUsername = i.User.Username
		save.CreateSave(PlayerSave)
	}
	questNumber = PlayerSave.Progress.Quest.QuestNumber
	if err != mongo.ErrNoDocuments && err != nil {
		log.Fatalf("Error generating quest: %v", err)
	}
	if h, ok := questsHandlers[questNumber]; ok {
		h(s, i)
	}
}
