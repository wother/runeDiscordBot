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

  // We check the message content and parse it
  if (msg.content === '!help') {
    let helpString = `
    The Rune Secrets Bot will draw runes for you from the elder Futhark.\n
    Commands are: \n
    **!cast** or **!castone** for a single rune casting\n
    **!castthree** for a three rune casting\n
    **!castfive** for a five rune casting (careful...)
    `;
    msg.channel.send(helpString);
  } else if (msg.content === '!cast') {
    let runeEmbed = runeToEmbed(randomRune(1), msg);

    msg.channel.send({ embed : runeEmbed });

  } else if (msg.content === '!castone') {
    let runeEmbed = runeToEmbed(randomRune(1), msg);

    msg.channel.send({ embed : runeEmbed });

  } else if (msg.content === '!castthree') {
    let runeArray = randomRune(3);
    runeArray.forEach(runeObj =>{

      msg.channel.send({ embed :  runeToEmbed(runeObj, msg) });

    });
  } else if (msg.content === '!castfive') {
    let runeArray = randomRune(5);
    runeArray.forEach(runeObj =>{

      msg.channel.send({ embed :  runeToEmbed(runeObj, msg) });
      
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