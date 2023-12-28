package save

import (
	"context"
	"fmt"
	config "tebu-discord/database/config"

	"gopkg.in/mgo.v2/bson"
)

// This does not update only one field, only full structs
func DeleteSave(discordId int64) {
	db := config.Collection
	filter := bson.M{
		"discord_id": discordId,
	}
	_, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Error updating save: %s\n", err)
		return
	}
}
