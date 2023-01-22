package eel

import (
	"github.com/bwmarrin/discordgo"
)

type EelConfig struct {
	DiscordToken   string
	DiscordAppId   string
	DiscordGuildId string
	DiscordSession *discordgo.Session
	TransposeToken string
}

func NewEelConfig(dtoken, appId, guildId, transpose string) *EelConfig {
	return &EelConfig{
		DiscordToken:   dtoken,
		DiscordAppId:   appId,
		DiscordGuildId: guildId,
		TransposeToken: transpose,
	}
}
