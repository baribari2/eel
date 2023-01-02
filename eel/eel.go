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

func NewEelConfig(discord, transpose string) *EelConfig {
	return &EelConfig{
		DiscordToken:   discord,
		TransposeToken: transpose,
	}
}
