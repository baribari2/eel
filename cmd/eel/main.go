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
	eel := eel.NewEelConfig(DISCORD_BOT_TOKEN, DISCORD_APPLICATION_ID, DISCORD_GUILD_ID, TRANSPOSE_API_KEY)

	err := start(eel)
	if err != nil {
		log.Fatal(err)
	}

	defer eel.DiscordSession.Close()
}

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
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "query",
						Description: "The SQL query to execute",
						Required:    true,
					},
				},
			},
		}
	)

	handlers := map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
		"query": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.ApplicationCommandData().Options[0].StringValue()

			res, err := transpose.ExecuteQuery(data, cfg)
			if err != nil {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("❌ Error executing query: %v", err.Error()),
					},
				})
			}

			r := fmt.Sprint(res.Results.([]interface{}))
			if len(r) < 2000 {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("✅ Query executed successfully. Results: %v", fmt.Sprint(res.Results)),
					},
				})
			} else {
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "✅ Query executed successfully. Results too large to display.",
					},
				})
			}

			if err != nil {
				log.Printf("\x1b[31m%s\x1b[0m%v", "❌ Error responding to interaction: ", err.Error())
			}
		},
	}

	err := discord.NewSession(cfg)
	if err != nil {
		return err
	}
	log.Printf("\x1b[32m%s\x1b[0m", "✅ Eel started...")

	cfg.DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	log.Printf("\x1b[32m%s\x1b[0m", "✅ Handlers added...")

	err = discord.RegisterCommands(commands, cfg)
	if err != nil {
		return err
	}
	log.Printf("\x1b[32m%s\x1b[0m", "✅ Commands registered...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("\x1b[33m%s\x1b[0m", "📣 Press Ctrl+C to exit")
	<-stop
	log.Printf("\x1b[33m%s\x1b[0m", "⚠️ Interrupt detected. Exiting...")

	cfg.DiscordSession.Close()

	return nil
}
