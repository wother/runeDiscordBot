// Run dotenv
require('dotenv').config();

// Import libraries
const { Client, MessageEmbed } = require('discord.js');
const client = new Client();

// workers imports
const randomRune = require('./workers/runes.js');

// Event listener when a user connected to the server.
client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

// Event listener when a user sends a message in the chat.
client.on('message', msg => {

  // We check the message content and looks for the word "ping", so we can have the bot respond "pong"
  if (msg.content === '!cast') {
    msg.reply(runeToEmbed(randomRune(1)));
  } else if (msg.content === '!castone') {
    msg.reply(runeToEmbed(randomRune(1)));
  } else if (msg.content === '!castthree') {
    msg.reply(runeToEmbed(randomRune(3)));
  } else if (msg.content === '!castfive') {
    msg.reply(runeToEmbed(randomRune(5)));
  }

});

function runeToEmbed (runeObject) {
  
  const embed = new MessageEmbed()
  .setTitle(runeObject.name)
  .setImage(runeObject.imgURL)
  .setURL(runeObject.descURL);

  return embed;
}

// Initialize bot by connecting to the server
try {
  client.login(process.env.DISCORD_TOKEN);
} catch(err) {
  throw new Error(err)
}