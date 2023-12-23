package service

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type session struct {
	botToken string
}

type SessionService interface {
	StartSession() *discordgo.Session
}

func New(botToken string) SessionService {
	return &session{botToken: botToken}
}

func (s *session) StartSession() *discordgo.Session {
	session, err := discordgo.New("Bot " + s.botToken)
	if err != nil {
		log.Fatalf("Error starting a new process: %s", err)
	}
	return session
}
