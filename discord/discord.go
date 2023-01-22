package discord

import (
	"eel/eel"
	"fmt"

	"github.com/broothie/qst"
	"github.com/bwmarrin/discordgo"
)

func NewSession(cfg *eel.EelConfig) error {
	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return err
	}

	err = dg.Open()
	if err != nil {
		return err
	}

	cfg.DiscordSession = dg

	return nil
}

func Send(data string, cfg *eel.EelConfig) error {
	_, err := cfg.DiscordSession.ChannelMessageSend("CHANNEL_ID", data)
	if err != nil {
		return err
	}

	return nil
}

func RegisterCommands(commands []*discordgo.ApplicationCommand, cfg *eel.EelConfig) error {
	for _, c := range commands {
		_, err := qst.Post(
			fmt.Sprintf("https://discord.com/api/v10/applications/%s/guilds/%s/commands", cfg.DiscordAppId, cfg.DiscordGuildId),
			qst.Header("Authorization", fmt.Sprintf("Bot %s", cfg.DiscordToken)),
			qst.BodyJSON(
				map[string]interface{}{
					"name":        c.Name,
					"type":        1,
					"description": c.Description,
					"options":     c.Options,
				},
			),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
