package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	functions "tebu-discord/functions"
	helper "tebu-discord/helper/env"
	service "tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

var (
	mainBotToken = "BOT_TOKEN"
	testBotToken = "BOT_TOKEN_TEST"
)

func main() {
	services := service.NewSessionService(helper.GetEnvValue(mainBotToken))
	session := services.StartSession()
	session.AddHandler(functions.MessageCreate)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := session.Open()
	if err != nil {
		log.Fatal("Error opening connection. Error: ", err)
	}
	defer session.Close()
	fmt.Println("The bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
