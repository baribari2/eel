module eel

go 1.19

require (
	github.com/broothie/qst v0.0.6
	github.com/bwmarrin/discordgo v0.26.1
	github.com/joho/godotenv v1.4.0
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
)

replace github.com/bwmarrin/discordgo v0.26.1 => ../discordgo
