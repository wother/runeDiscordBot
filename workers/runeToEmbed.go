package workers

import (
	"github.com/bwmarrin/discordgo"
)

// RuneToMessage converts a rune to a Discord message.
func RuneToMessage(runeObj Rune) *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{runeToEmbed(runeObj)},
		Components: []discordgo.MessageComponent{
			componentRowWith(runeLinkButton(runeObj.Name, runeObj.DescURL)),
		},
	}
}

func runeToEmbed(runeObj Rune) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       runeObj.Name,
		URL:         runeObj.DescURL,
		Description: "You drew " + runeObj.Name,
		Image: &discordgo.MessageEmbedImage{
			URL: runeObj.ImgURL,
		},
	}
}

func componentRowWith(button discordgo.Button) discordgo.ActionsRow {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button},
	}
}

func runeLinkButton(title, url string) discordgo.Button {
	return discordgo.Button{
		Label: "Go to " + title + " info page.",
		Style: discordgo.LinkButton,
		URL:   url,
	}
}

func AllRunesLinks(futharkArray []string) *discordgo.MessageSend {
	output := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{},
	}
	runeNameMatrix := ChunkArray(futharkArray, 5)
	for _, nameArray := range runeNameMatrix {
		messageRow := discordgo.ActionsRow{}
		for _, runeName := range nameArray {
			messageRow.Components = append(messageRow.Components, runeNameButton(runeName))
		}
		output.Components = append(output.Components, messageRow)
	}
	return output
}

func runeNameButton(name string) discordgo.Button {
	return discordgo.Button{
		Label: name,
		Style: discordgo.LinkButton,
		URL:   genInfoLink(name),
	}
}