package gatherwood

import (
	"fmt"
	"strconv"
	"tebu-discord/database/models"

	"github.com/bwmarrin/discordgo"
)

var (
	points int
)

func GatherWoodButton(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	playerSave ...*models.PlayerSave,
) {
	points++
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "value: " + strconv.Itoa(points),
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{discordgo.Button{
					Label: "Gather some sticks",
					Style: discordgo.SuccessButton,
					Emoji: discordgo.ComponentEmoji{
						Name: "🪵",
					},
					CustomID: "gather_wood_button",
				}},
			}},
		},
	})
	if err != nil {
		fmt.Printf("Error creating increment button: %v \n", err)
	}
}
