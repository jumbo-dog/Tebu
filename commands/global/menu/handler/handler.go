package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func MenuHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			CustomID: "main_menu_options",
			Title:    "Menu:",
			Flags:    discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Play game",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "main_menu_play",
							Emoji: discordgo.ComponentEmoji{
								Name: "🎮",
							},
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Github",
							Style:    discordgo.PrimaryButton,
							Disabled: false,
							CustomID: "main_menu_github",
							Emoji: discordgo.ComponentEmoji{
								Name: "💻",
							},
						},
					},
				},
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Global ranking",
							Style:    discordgo.LinkButton,
							URL:      "https://www.youtube.com/watch?v=oiNPgJmtzVI",
							Disabled: false,
							Emoji: discordgo.ComponentEmoji{
								Name: "🏆",
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Menu error: %s ", err)
	}
}
