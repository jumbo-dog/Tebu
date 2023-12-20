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

func NewSessionService(botToken string) SessionService {
	return &session{botToken: botToken}
}
func (s *session) StartSession() *discordgo.Session {
	session, err := discordgo.New("Bot " + s.botToken)
	if err != nil {
		log.Fatal("Error starting a new process", err)
	}
	return session
}
