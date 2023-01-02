package main

import (
	"eel/discord"
	"eel/eel"
	"eel/transpose"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	eel := eel.NewEelConfig(DISCORD_BOT_TOKEN, TRANSPOSE_API_KEY)

	err := start(eel)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Add guild id
func start(cfg *eel.EelConfig) error {
	var (
		dmPerm                  = false
		defaultMemberPerm int64 = discordgo.PermissionManageServer

		commands = []*discordgo.ApplicationCommand{
			{
				Name:                     "query",
				Description:              "Execute a SQL query on ethereum data",
				DefaultMemberPermissions: &defaultMemberPerm,
				DMPermission:             &dmPerm,
			},
		}
	)

	defer cfg.DiscordSession.Close()

	handlers := map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		"eel-query": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.ApplicationCommandData().Options[0].StringValue()

			res, err := transpose.ExecuteQuery(data, cfg)
			if err != nil {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Error executing query: %v", err.Error()),
					},
				})
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Query executed successfully. Results: %v", res.Results),
				},
			})
		},
	}

	log.Printf("\x1b[32m%s\x1b[0m", "Starting Eel...")
	err := discord.NewSession(cfg)
	if err != nil {
		return err
	}

	log.Printf("\x1b[32m%s\x1b[0m", "Adding handlers...")
	cfg.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	log.Printf("\x1b[32m%s\x1b[0m", "Registering commands...")
	err = discord.RegisterCommands(commands, cfg)
	if err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("\x1b[33m%s\x1b[0m", "Press Ctrl+C to exit")
	<-stop
	log.Printf("\x1b[33m%s\x1b[0m", "Interrupt detected. Exiting...")

	return nil
}
