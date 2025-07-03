# Rune Draw Bot for Discord #

v3 is live and stable. Need to circle back around on the API and docs, probably reformat the help output to be a bit more polished... dunno.

This is a Discord bot that draws runes from the Elder Futhark and displays them in a channel.

### Usage: ###

Add the bot to your server with [this link](https://discord.com/api/oauth2/authorize?client_id=714237973486633031&permissions=18432&scope=bot).

You will also need to put your token into the `env` file in the root of the project,
rename it to `.env` to ensure it is covered under the `.gitignore` and you don't
publicise the token.

- `!help` for printing help text and command list.
- `!draw [number]`  draw and display this many runes, default is one
- `!list` gives you a referance for rune names, alphabetically.
- `!lookup` [runeName] fetches the meaning for a rune from the Rune Secrets Site

## LICENCE ##

This code repository is covered under the MIT Licence [available here](./LICENCE)
