// Run dotenv
require('dotenv').config();

// Import libraries
const Discord = require('discord.js');
const client = new Discord.Client();

// workers imports
// const rantInt = require('./workers/randomizer.js')

// Event listener when a user connected to the server.
client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
});

// Event listener when a user sends a message in the chat.
client.on('message', msg => {

  // We check the message content and looks for the word "ping", so we can have the bot respond "pong"
  var msgContentString = new String(msg.content);
  if (msgContentString.includes('ping')) {
    msg.reply('pong ' + msg.content);
    // TODO: Rune logic calls go here.
  } 

});

// Initialize bot by connecting to the server
try {
  client.login(process.env.DISCORD_TOKEN);
} catch(err) {
  throw new Error(err)
}