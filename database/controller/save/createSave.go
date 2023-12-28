package save

import (
	"context"
	"fmt"
	config "tebu-discord/database/config"
	"tebu-discord/database/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ()

func CreateSave(information *models.PlayerSave) {
	db := config.Collection
	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{"discord_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		fmt.Printf("Error creating unique index: %s\n", err)
		return
	}

	_, err = db.InsertOne(context.Background(), information)
	if err != nil {
		fmt.Printf("Error creating new save: %s\n", err)
		return
	}
}
