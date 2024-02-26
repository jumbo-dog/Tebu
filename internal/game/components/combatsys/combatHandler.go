package combatsys

import (
	"log"
	"math/rand"
	"strconv"
	"tebu-discord/database/controller/save"
	"tebu-discord/internal/models"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	florestEnemys = []models.Enemy{
		{
			Name:        "Paegorn's",
			Health:      10,
			Description: "big man little hands",
		},
		{
			Name:        "Discord moderator (corrupted)",
			Health:      11,
			Description: "lalala",
		},
	}
	dungeonEnemys = []models.Enemy{
		{
			Name:        "Jonnark",
			Health:      20,
			Description: "big man little hand's",
		},
		{
			Name:        "Donald Reginald Trump's",
			Health:      11,
			Description: "lalala",
		},
		{
			Name:        "Dumb son of bitch's",
			Health:      15,
			Description: "lalala",
		},
		{
			Name:        "heavy from tf2??",
			Health:      15,
			Description: "lalala",
		},
	}
	playerHealth     int = 0
	enemyHealth      int = 0
	enemyName        string
	enemyDescription string
	goBackId         string
	buttonIdLoose    string
	damage           = 1
)

func HandleCombat(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	if playerHealth <= 0 || enemyHealth <= 0 {
		log.Println("Error starting combat: invalid player or enemy health")
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	msg, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Flags:   discordgo.MessageFlagsEphemeral,
		Content: "You've found a wild " + enemyName,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Attack",
						Style:    discordgo.DangerButton,
						CustomID: "attack_button",
						Emoji: discordgo.ComponentEmoji{
							Name: "⚔️",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("Error creating follow up message: %v\n", err)
		return
	}

	for {
		paragraph := "Your health: " + strconv.Itoa(
			playerHealth,
		) + "\n" + enemyName + " health: " + strconv.Itoa(
			enemyHealth,
		) + "\n"

		s.FollowupMessageEdit(i.Interaction, msg.ID, &discordgo.WebhookEdit{
			Content: &paragraph,
		})

		playerHealth -= 1

		if playerHealth <= 0 {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "You lost to " + enemyName,
				Flags:   discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					&discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label: "Go back",
								Style: discordgo.SuccessButton,
								Emoji: discordgo.ComponentEmoji{
									Name: "🗿",
								},
								CustomID: buttonIdLoose,
							},
						},
					},
				},
			})
			break
		}
		if enemyHealth <= 0 {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "You won",
				Flags:   discordgo.MessageFlagsEphemeral,
				Components: []discordgo.MessageComponent{
					&discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label: "Go back",
								Style: discordgo.SuccessButton,
								Emoji: discordgo.ComponentEmoji{
									Name: "🗿",
								},
								CustomID: goBackId,
							},
						},
					},
				},
			})
			break
		}

		time.Sleep(time.Second * 1)
	}
}

const (
	FOREST_2 int = iota
	DUNGEON
)

func GetCombatStats(location int, playerID, buttonId, buttonIdLost string) {
	playerHealth = 10
	playerSave, err := save.GetSave(playerID)
	if err != nil {
		log.Println("Error: Couldn't get player save: ", err)
		return
	}
	playerItems := playerSave.Items

	if playerItems["armor"] > 0 {
		playerHealth += 5
	}
	if playerItems["wooden_spear"] > 0 {
		damage = 3
	}
	var enemy models.Enemy
	switch location {
	case FOREST_2:
		enemy = chooseEnemy(florestEnemys)
	case DUNGEON:
		enemy = chooseEnemy(dungeonEnemys)
	}
	enemyHealth = enemy.Health
	enemyName = enemy.Name
	enemyDescription = enemy.Description
	goBackId = buttonId
	buttonIdLoose = buttonIdLost
}

func chooseEnemy(enemys []models.Enemy) models.Enemy {
	randomIndex := rand.Intn(len(enemys))
	randomItem := enemys[randomIndex]
	return randomItem
}

func HandleAttacks(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})
	enemyHealth -= damage
}
