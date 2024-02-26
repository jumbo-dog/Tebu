package save

import (
	"context"
	"fmt"
	"log"
	config "tebu-discord/database/config"
	"tebu-discord/database/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// This does not update only one field, only full structs
func GetSave(discordId string) (*models.PlayerSave, error) {
	result := &models.PlayerSave{
		DiscordId: discordId,
		Progress: &models.Progress{
			Planet: &models.Planet{},
			Quest:  &models.Quest{},
		},
	}
	if discordId == "" {
		fmt.Printf("Discord id is obligatory:\n")
		return result, nil
	}
	db := config.Collection
	filter := bson.M{"discord_id": discordId}

	err := db.FindOne(context.Background(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Save not found: %s\n", err)
		return result, err
	}
	if err != nil {
		log.Fatal("Error obtaining the save", err)
	}
	return result, nil
}
