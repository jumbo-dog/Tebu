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
	questNumber = PlayerSave.Progress.Quest.QuestNumber
	if err == mongo.ErrNoDocuments {
		newSave := &models.PlayerSave{
			DiscordId:    i.User.ID,
			LastUsername: i.User.Username,
			Progress: &models.Progress{
				Quest: &models.Quest{
					QuestNumber: 0,
				},
			},
		}
		questNumber = 0
		save.CreateSave(newSave)
	}
	if err != nil && err != mongo.ErrNoDocuments {
		log.Fatalf("Error generating quest: %v", err)
	}
	if h, ok := questsHandlers[questNumber]; ok {
		h(s, i, PlayerSave)
	}
}
