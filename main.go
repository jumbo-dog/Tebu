package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	global "tebu-discord/commands/global/menu"
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
	services := service.New(helper.GetEnvValue(mainBotToken))
	global := global.New(services)

	session := services.StartSession()
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	defer session.Close()

	session.AddHandler(functions.MessageCreate)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := session.Open()
	if err != nil {
		log.Fatal("Error opening connection. Error: ", err)
	}

	session.AddHandler(global.StartMenu)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
