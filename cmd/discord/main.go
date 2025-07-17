package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"runeDiscordBot/internal/workers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}

	// Discord bot mode - require token
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
	// Register the interaction handler for button clicks
	dg.AddHandler(interactionCreate)

	// Set up intents to receive message events and content
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("Error: Port is already in use. Is the bot already running?")
			return
		}
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

	// Add debug logging
	fmt.Printf("Received message: %s\n", m.Content)

	// Parse the message
	response := workers.ParseMessage(m.Content)

	// If there is a response, send it
	if response != nil {
		fmt.Printf("Response type: %s\n", response["type"])

		switch response["type"] {
		case "text":
			_, err := s.ChannelMessageSend(m.ChannelID, response["content"].(string))
			if err != nil {
				fmt.Printf("Error sending text message: %v\n", err)
			}
		case "embed":
			_, err := s.ChannelMessageSendComplex(m.ChannelID, workers.RuneToMessage(response["content"].(workers.Rune)))
			if err != nil {
				fmt.Printf("Error sending embed message: %v\n", err)
			}
		case "info":
			// Handle info type - this was missing!
			_, err := s.ChannelMessageSendComplex(m.ChannelID, workers.RuneToMessage(response["content"].(workers.Rune)))
			if err != nil {
				fmt.Printf("Error sending info message: %v\n", err)
			}
		case "runeArray":
			// Combine multiple runes into a single message
			runeArray := response["content"].([]workers.Rune)
			combinedMessage := workers.RuneArrayToMessage(runeArray)
			_, err := s.ChannelMessageSendComplex(m.ChannelID, combinedMessage)
			if err != nil {
				fmt.Printf("Error sending rune array message: %v\n", err)
			}
		case "allRunesLinks":
			_, err := s.ChannelMessageSendComplex(m.ChannelID, response["content"].(*discordgo.MessageSend))
			if err != nil {
				fmt.Printf("Error sending all runes links message: %v\n", err)
			}
		default:
			fmt.Printf("Unknown response type: %s\n", response["type"])
		}
	} else {
		fmt.Printf("No response generated for message: %s\n", m.Content)
	}
}

// Handle button interactions
func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Only handle button interactions
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	customID := i.MessageComponentData().CustomID
	fmt.Printf("Button interaction received: %s\n", customID)

	// Parse the custom ID to get action and rune name
	// Format: "info:[runename]" or "back:[runename]"
	parts := strings.Split(customID, ":")
	if len(parts) != 2 {
		return
	}

	action := parts[0]
	runeName := parts[1]

	switch action {
	case "info":
		// Show info version with info png
		infoMessage := workers.RuneToInfoMessage(workers.Rune{
			Name:    runeName,
			ImgFile: fmt.Sprintf("media/info/rune-info-%s.png", runeName),
			DescURL: fmt.Sprintf("https://example.com/rune-meanings/%s", runeName),
		})

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:      infoMessage.Embeds,
				Components:  infoMessage.Components,
				Files:       infoMessage.Files,
				Attachments: &[]*discordgo.MessageAttachment{}, // Clear old attachments
			},
		})
		if err != nil {
			fmt.Printf("Error responding to info interaction: %v\n", err)
		}

	case "back":
		// Show original rune version
		originalMessage := workers.RuneToMessage(workers.Rune{
			Name:    runeName,
			ImgFile: fmt.Sprintf("media/solo/%s.png", runeName),
			DescURL: fmt.Sprintf("https://example.com/rune-meanings/%s", runeName),
		})

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:      originalMessage.Embeds,
				Components:  originalMessage.Components,
				Files:       originalMessage.Files,
				Attachments: &[]*discordgo.MessageAttachment{}, // Clear old attachments
			},
		})
		if err != nil {
			fmt.Printf("Error responding to back interaction: %v\n", err)
		}
	}
}