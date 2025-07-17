package workers

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// RuneToMessage converts a rune to a Discord message.
func RuneToMessage(runeObj Rune) *discordgo.MessageSend {
	fmt.Printf("Creating message for rune: %s\n", runeObj.Name)

	// Try to load the image file
	img, err := os.Open(runeObj.ImgFile)
	if err != nil {
		fmt.Printf("Could not open image file %s: %v. Using text-only embed.\n", runeObj.ImgFile, err)
		// Return message without image
		return &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{runeToEmbedNoImage(runeObj)},
			Components: []discordgo.MessageComponent{
				componentRowWith(runeLinkButton(runeObj.Name, runeObj.DescURL)),
			},
		}
	}
	defer img.Close()

	// Read image into buffer
	imgBuff := new(bytes.Buffer)
	_, err = imgBuff.ReadFrom(img)
	if err != nil {
		fmt.Printf("Could not read image file %s: %v. Using text-only embed.\n", runeObj.ImgFile, err)
		// Return message without image
		return &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{runeToEmbedNoImage(runeObj)},
			Components: []discordgo.MessageComponent{
				componentRowWith(runeLinkButton(runeObj.Name, runeObj.DescURL)),
			},
		}
	}

	fmt.Printf("Successfully loaded image for %s\n", runeObj.Name)
	// Return message with image
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{runeToEmbed(runeObj)},
		Components: []discordgo.MessageComponent{
			componentRowWith(runeLinkButton(runeObj.Name, runeObj.DescURL)),
		},
		Files: []*discordgo.File{
			{
				Name:   strings.Split(runeObj.ImgFile, "/")[2], // Just the filename
				Reader: imgBuff,
			},
		},
	}
}

func runeToEmbed(runeObj Rune) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name,
		URL:         runeObj.DescURL,
		Description: "You drew " + runeObj.Name,
		Image: &discordgo.MessageEmbedImage{
			URL: "attachment://" + strings.Split(runeObj.ImgFile, "/")[2],
		},
	}
}

func runeToEmbedNoImage(runeObj Rune) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name,
		URL:         runeObj.DescURL,
		Description: "You drew " + runeObj.Name + " (image not available)",
		Color:       0x9966cc, // Purple color for fallback
	}
}

func componentRowWith(button discordgo.Button) discordgo.ActionsRow {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}
}

func runeLinkButton(title, url string) discordgo.Button {
	return discordgo.Button{
		Label:    "Show " + title + " info",
		Style:    discordgo.PrimaryButton,
		CustomID: "info:" + title,
	}
}



// RuneToInfoMessage converts a rune to a Discord message with info gif
func RuneToInfoMessage(runeObj Rune) *discordgo.MessageSend {
	fmt.Printf("Creating info message for rune: %s\n", runeObj.Name)

	// Try to load the info gif file
	img, err := os.Open(runeObj.ImgFile)
	if err != nil {
		fmt.Printf("Could not open info file %s: %v. Using text-only embed.\n", runeObj.ImgFile, err)
		// Return message without image but with back button
		return &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{runeToInfoEmbedNoImage(runeObj)},
			Components: []discordgo.MessageComponent{
				componentRowWith(runeBackButton(runeObj.Name)),
			},
		}
	}
	defer img.Close()

	// Read image into buffer
	imgBuff := new(bytes.Buffer)
	_, err = imgBuff.ReadFrom(img)
	if err != nil {
		fmt.Printf("Could not read info file %s: %v. Using text-only embed.\n", runeObj.ImgFile, err)
		// Return message without image but with back button
		return &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{runeToInfoEmbedNoImage(runeObj)},
			Components: []discordgo.MessageComponent{
				componentRowWith(runeBackButton(runeObj.Name)),
			},
		}
	}

	fmt.Printf("Successfully loaded info for %s\n", runeObj.Name)
	// Return message with info gif
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{runeToInfoEmbed(runeObj)},
		Components: []discordgo.MessageComponent{
			componentRowWith(runeBackButton(runeObj.Name)),
		},
		Files: []*discordgo.File{
			{
				Name:   strings.Split(runeObj.ImgFile, "/")[2], // Just the filename
				Reader: imgBuff,
			},
		},
	}
}

func runeToInfoEmbed(runeObj Rune) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name + " - Detailed Information",
		Description: "Detailed meaning and interpretation of " + runeObj.Name,
		Color:       0x00ff00, // Green color for info
		Image: &discordgo.MessageEmbedImage{
			URL: "attachment://" + strings.Split(runeObj.ImgFile, "/")[2],
		},
	}
}

func runeToInfoEmbedNoImage(runeObj Rune) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name + " - Detailed Information",
		Description: "Detailed meaning and interpretation of " + runeObj.Name + " (info gif not available)",
		Color:       0x00ff00, // Green color for info
	}
}

func runeBackButton(runeName string) discordgo.Button {
	return discordgo.Button{
		Label:    "← Back to " + runeName,
		Style:    discordgo.SecondaryButton,
		CustomID: "back:" + runeName,
	}
}

// RuneArrayToMessage converts multiple runes to a single Discord message with multiple embeds
func RuneArrayToMessage(runes []Rune) *discordgo.MessageSend {
	fmt.Printf("Creating combined mystery message for %d runes\n", len(runes))

	var embeds []*discordgo.MessageEmbed
	var files []*discordgo.File
	var components []discordgo.MessageComponent

	// Create a row of buttons for each rune
	var buttonRows []discordgo.MessageComponent
	currentRow := discordgo.ActionsRow{}
	buttonCount := 0

	for _, runeObj := range runes {
		// Use small mysterious images for array casting
		smallImgPath := genSmallImgPath(runeObj.Name)

		// Try to load the small image file
		img, err := os.Open(smallImgPath)
		if err != nil {
			fmt.Printf("Could not open small image file %s: %v. Using text-only embed.\n", smallImgPath, err)
			// Add text-only embed
			embeds = append(embeds, runeToEmbedNoImage(runeObj))
		} else {
			// Read image into buffer
			imgBuff := new(bytes.Buffer)
			_, err = imgBuff.ReadFrom(img)
			img.Close()

			if err != nil {
				fmt.Printf("Could not read small image file %s: %v. Using text-only embed.\n", smallImgPath, err)
				// Add text-only embed
				embeds = append(embeds, runeToEmbedNoImage(runeObj))
			} else {
				// Add embed with small mysterious image
				embeds = append(embeds, runeToMysteryEmbed(runeObj))
				// Add file attachment
				files = append(files, &discordgo.File{
					Name:   strings.Split(smallImgPath, "/")[len(strings.Split(smallImgPath, "/"))-1],
					Reader: imgBuff,
				})
			}
		}

		// Add info button for this rune
		infoButton := discordgo.Button{
			Label:    runeObj.Name + " info",
			Style:    discordgo.PrimaryButton,
			CustomID: "info:" + runeObj.Name,
		}

		currentRow.Components = append(currentRow.Components, infoButton)
		buttonCount++

		// Discord allows max 5 buttons per row
		if buttonCount >= 5 {
			buttonRows = append(buttonRows, currentRow)
			currentRow = discordgo.ActionsRow{}
			buttonCount = 0
		}
	}

	// Add any remaining buttons
	if buttonCount > 0 {
		buttonRows = append(buttonRows, currentRow)
	}

	components = buttonRows

	return &discordgo.MessageSend{
		Content:    fmt.Sprintf("🌙 **%d-Rune Mystical Casting** 🌙\n*The runes whisper their secrets...*", len(runes)),
		Embeds:     embeds,
		Components: components,
		Files:      files,
	}
}

// RuneToMysteryMessage creates a mysterious small version of a rune
func RuneToMysteryMessage(runeObj Rune) *discordgo.MessageSend {
	fmt.Printf("Creating mystery message for rune: %s\n", runeObj.Name)

	// Use small image path for mystery
	smallImgPath := genSmallImgPath(runeObj.Name)

	// Try to load the small image file
	img, err := os.Open(smallImgPath)
	if err != nil {
		fmt.Printf("Could not open small image file %s: %v. Using regular size.\n", smallImgPath, err)
		// Fallback to regular message
		return RuneToMessage(runeObj)
	}
	defer img.Close()

	// Read image into buffer
	imgBuff := new(bytes.Buffer)
	_, err = imgBuff.ReadFrom(img)
	if err != nil {
		fmt.Printf("Could not read small image file %s: %v. Using regular size.\n", smallImgPath, err)
		// Fallback to regular message
		return RuneToMessage(runeObj)
	}

	fmt.Printf("Successfully loaded small image for %s\n", runeObj.Name)
	filename := strings.Split(smallImgPath, "/")[len(strings.Split(smallImgPath, "/"))-1]
	fmt.Printf("Small image filename: %s\n", filename)
	// Return message with small mysterious image
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{runeToMysteryEmbed(runeObj)},
		Components: []discordgo.MessageComponent{
			componentRowWith(runeLinkButton(runeObj.Name, runeObj.DescURL)),
		},
		Files: []*discordgo.File{
			{
				Name:   strings.Split(smallImgPath, "/")[len(strings.Split(smallImgPath, "/"))-1], // Get the last part (filename)
				Reader: imgBuff,
			},
		},
	}
}

func runeToMysteryEmbed(runeObj Rune) *discordgo.MessageEmbed {
	smallImgPath := genSmallImgPath(runeObj.Name)
	filename := strings.Split(smallImgPath, "/")[len(strings.Split(smallImgPath, "/"))-1]
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name,
		Description: "A rune emerges from the aether...",
		Color:       0x9966cc, // Purple color for mystery
		Image: &discordgo.MessageEmbedImage{
			URL: "attachment://" + filename,
		},
	}
}
