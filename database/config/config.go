package controller

import (
	"context"
	"fmt"
	"log"
	helper "tebu-discord/internal/helper/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	Collection *mongo.Collection
	err        error
)

func ConnectDatabase() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://" + helper.GetEnvValue("DATABASE_LOGIN", "../../.env") + ":" + helper.GetEnvValue("DATABASE_PASSWORD", "../../.env") + "@cluster0.iwwx7zx.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Error connecting to mongo: ", err)
	}
	Collection = client.Database("tebu").Collection("player-progress")
	fmt.Println("Connected to MongoDB!")
}

func DisconnectDatabase() {
	fmt.Println("Disconnecting MongoDB")
	if err = client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error disconnecting from mongo: ", err)
	}
}
