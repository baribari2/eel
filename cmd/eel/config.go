package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	TRANSPOSE_API_KEY      string
	DISCORD_BOT_TOKEN      string
	DISCORD_APPLICATION_ID string
	DISCORD_GUILD_ID       string
)

func getEnv(env string) string {
	env, exists := os.LookupEnv(env)
	if !exists {
		log.Fatal("Environment variable not set: " + env)
	}

	return env
}

func init() {
	godotenv.Load("../../.env")

	TRANSPOSE_API_KEY = getEnv("TRANSPOSE_API_KEY")
	DISCORD_BOT_TOKEN = getEnv("DISCORD_BOT_TOKEN")
	DISCORD_APPLICATION_ID = getEnv("DISCORD_APPLICATION_ID")
	DISCORD_GUILD_ID = getEnv("DISCORD_GUILD_ID")
}
