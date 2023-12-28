package save

import (
	"context"
	"fmt"
	config "tebu-discord/database/config"
	"tebu-discord/database/models"

	"gopkg.in/mgo.v2/bson"
)

// This does not update only one field, only full structs
func UpdateSave(information *models.PlayerSave) {
	if information.DiscordId == "" {
		fmt.Printf("Discord id is obligatory:\n")
		return
	}
	db := config.Collection
	filter := bson.M{
		"discord_id": information.DiscordId,
	}
	update := bson.M{
		"$set": information,
	}
	_, err := db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("Error updating save: %s\n", err)
		return
	}
}
