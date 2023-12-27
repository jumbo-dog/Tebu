package save

import (
	"context"
	"fmt"
	config "tebu-discord/database/config"
	"tebu-discord/database/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// This does not update only one field, only full structs
func GetSave(discordId int64) *models.PlayerSave {
	result := &models.PlayerSave{
		DiscordId: discordId,
	}
	if discordId == 0 {
		fmt.Printf("Discord id is obligatory:")
		return result
	}
	db := config.Collection
	filter := bson.M{"discord_id": discordId}

	err := db.FindOne(context.Background(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Save not found: %s", err)
		return result
	}
	if err != nil {
		fmt.Printf("Error obtaining the save: %s", err)
		return result
	}
	return result
}
