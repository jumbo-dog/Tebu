package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	direct "tebu-discord/commands/direct/game/gatherButton"
	menu "tebu-discord/commands/global/menu"
	helper "tebu-discord/helper/env"
	service "tebu-discord/service"

	"github.com/bwmarrin/discordgo"
)

var (
	mainBotToken = "BOT_TOKEN"
	testBotToken = "BOT_TOKEN_TEST"
	created      = false
)

func main() {
	services := service.New(helper.GetEnvValue(mainBotToken))
	session := services.StartSession()
	menu := menu.New(services)
	if created == false {
		direct := direct.New(services)
		session.AddHandler(direct.GatherButton)
	}
	session.AddHandler(menu.StartMenu)

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	defer session.Close()

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := session.Open()
	if err != nil {
		log.Fatal("Error opening connection. Error: ", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
