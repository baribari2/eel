package discord

import (
	"eel/eel"

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
	for _, command := range commands {
		_, err := cfg.DiscordSession.ApplicationCommandCreate(cfg.DiscordAppId, "", command)
		if err != nil {
			return err
		}

	}

	return nil
}
