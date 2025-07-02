package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"runeDiscordBot/api"
	"runeDiscordBot/workers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}
	// Create a new Discord session using the provided bot token.
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		token = os.Getenv("TEST_BOT_TOKEN") // Fallback for testing
	}
	if token == "" {
		fmt.Println("Error: No Discord token found. Please set DISCORD_TOKEN or TEST_BOT_TOKEN environment variable.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Run the API in a goroutine
	go api.SetupAPI()

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Parse the message
	response := workers.ParseMessage(m.Content)

	// If there is a response, send it
	if response != nil {
		switch response["type"] {
		case "text":
			s.ChannelMessageSend(m.ChannelID, response["content"].(string))
		case "embed":
			s.ChannelMessageSendComplex(m.ChannelID, workers.RuneToMessage(response["content"].(workers.Rune)))
		case "runeArray":
			for _, runeObj := range response["content"].([]workers.Rune) {
				s.ChannelMessageSendComplex(m.ChannelID, workers.RuneToMessage(runeObj))
			}
		case "allRunesLinks":
			s.ChannelMessageSendComplex(m.ChannelID, response["content"].(*discordgo.MessageSend))
		}
	}
}
