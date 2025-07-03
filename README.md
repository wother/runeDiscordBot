# Rune Draw Bot for Discord #

v3.0.0 is live and stable! This is a complete Go rewrite with interactive buttons, multi-rune casting, and rare variant images.

This is a Discord bot that draws runes from the Elder Futhark and displays them in a channel with beautiful images and interactive info buttons.

## 🚀 Quick Setup

### Automated Setup (Recommended)
Run the setup script to automatically install Go, configure the bot, and build the binary:

```bash
./setup.sh
```

This will:
- Check/install Go (if needed)
- Prompt for your Discord Bot Token
- Generate a `.env` configuration file
- Build the Linux binary for deployment

### Manual Setup
1. Install Go 1.19+ from https://golang.org/dl/
2. Copy `env-example` to `.env` and add your Discord Bot Token
3. Build the project: `go build -o runeDiscordBot-linux main.go`
4. Run: `./runeDiscordBot-linux`

## 📖 Usage

Add the bot to your server with [this link](https://discord.com/api/oauth2/authorize?client_id=714237973486633031&permissions=18432&scope=bot).

### Commands:
- `!help` - Show help text and command list
- `!rune` or `!draw [number]` - Draw and display runes (default: 1)
- `!cast [number]` - Mystical multi-rune casting
- `!info [runeName]` - Get detailed information about a specific rune
- `!list` - Reference list of all rune names

### Features:
- 🎯 Interactive info buttons - click to swap between rune image and detailed info
- 🌙 Multi-rune casting in single messages with individual info buttons
- ✨ 5% chance for rare "dreams" variant images
- 🔮 Beautiful Elder Futhark rune images with mystical theming

## 📁 Deployment Files

For deployment, you need:
- `runeDiscordBot-linux` (the binary)
- `.env` (your configuration)
- `media/` (the rune images folder)

## LICENCE

This code repository is covered under the MIT Licence [available here](./LICENSE)
