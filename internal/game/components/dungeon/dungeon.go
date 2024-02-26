package dungeon

import (
	"fmt"
	"tebu-discord/internal/game/components/combatsys"

	"github.com/bwmarrin/discordgo"
)

var (
	paragraph              = "=================================================================\n===============================================\n ENTERING DUNGEON \nnYou step into the damp dungeon, the air thick with anticipation and the echo of distant drips. Shadows dance across ancient stone walls as your torch flickers, illuminating the unknown ahead."
	buttonId        string = "explore_dungeon"
	dungeonProgress int
	disableExplore  bool
)

func StartDungeon(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {

	combatsys.GetCombatStats(combatsys.DUNGEON, i.User.ID, "explore_dungeon", "goto_camp")
	dungeonProgression(i)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: paragraph,
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label: "Explore",
							Style: discordgo.SuccessButton,
							Emoji: discordgo.ComponentEmoji{
								Name: "🧭",
							},
							Disabled: disableExplore,
							CustomID: buttonId,
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("Error creating increment button: %v \n", err)
	}
}

func dungeonProgression(i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "explore" {
		dungeonProgress += 1
		if dungeonProgress == 2 {
			paragraph = "As you venture deeper into the dungeon, a menacing growl fills the air, sending shivers down your spine. Your heart races as you brace yourself to confront the lurking monster"
		}
		if dungeonProgress == 3 {
			buttonId = "init_attack"
		}
		if dungeonProgress == 4 {
		}
		if dungeonProgress == 5 {
		}
		if dungeonProgress == 6 {
		}
	}

}
