// Run dotenv
require('dotenv').config();

// Import libraries
const { Client, MessageEmbed } = require('discord.js');
const client = new Client();

// workers imports
const parseMessage = require("./workers/commandStringParse.js");

// Event listener when bot connects to the server.
client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
  client.user.setPresence({ game: { name: '!help for commands' , type: 'LISTENING' }, status: 'active' })
    .then(console.log("We are ready for input, Dave."))
    .catch(console.error);
});

// Event listener when a user sends a message in the chat.
client.on('message', msg => {

  // We check the message content and parse it
  
  let parsedMessage = parseMessage(msg.content);
  if (parsedMessage && parsedMessage.type === "text") {
    msg.channel.send(parsedMessage.content);
  } else if (parsedMessage && parsedMessage.type === "embed") {
    let runeEmbed = runeToEmbed(parsedMessage.content, msg);
    msg.channel.send({ embed : runeEmbed});
  } else if (parsedMessage && parsedMessage.type === "runeArray") {
    parsedMessage.content.forEach(runeObj => {
      msg.channel.send({"embed": runeToEmbed(runeObj, msg)});
    });
  }
});

function runeToEmbed (runeObject, inputMessage) {

  const embed = new MessageEmbed(inputMessage, {
    "title" : runeObject.name,
    "url"   : runeObject.descURL,
    "image" : {
      "url"     : runeObject.imgURL
    }
  });

  return embed;
};

// Initialize bot by connecting to the server
try {
  client.login(process.env.DISCORD_TOKEN);
} catch(err) {
  throw new Error(err)
}