package save

import (
	"context"
	"fmt"
	config "tebu-discord/database/config"
	"tebu-discord/database/models"
	"tebu-discord/pkg/reflection"

	"gopkg.in/mgo.v2/bson"
)

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
	lastSave, err := GetSave(information.DiscordId)
	if err != nil {
		fmt.Println("Failed to retrieve information")
		return
	}
	reflection.AttributeValues(information, lastSave)
	reflection.InitializeStructIfNil(information, lastSave.Progress, "Progress")
	reflection.AttributeValues(information.Progress, lastSave.Progress)

	reflection.InitializeStructIfNil(information.Progress, lastSave.Progress.Planet, "Planet")
	reflection.AttributeValues(information.Progress.Planet, lastSave.Progress.Planet)

	reflection.InitializeStructIfNil(information.Progress, lastSave.Progress.Quest, "Quest")
	reflection.AttributeValues(information.Progress.Quest, lastSave.Progress.Quest)

	_, err = db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("Error updating save: %s\n", err)
		return
	}
}
