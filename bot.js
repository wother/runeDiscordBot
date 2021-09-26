import {} from 'dotenv/config.js';
// Import libraries
import { Client, Presence, Intents } from 'discord.js';
const client = new Client(
  {
    intents : [
      Intents.FLAGS.GUILDS,
      Intents.FLAGS.GUILD_MESSAGES,
      Intents.FLAGS.GUILD_PRESENCES
    ]
  }
);

// workers imports
import { allRunesLinks, runeToMessage } from "./workers/runeToEmbed.js";
import { randomRune, getFutharkArray } from "./workers/runes.js";
import { parseMessage } from "./workers/commandStringParse.js";

// Event listener when a user connected to the server.
client.on('ready', () => {
  console.log(`Logged in as ${client.user.tag}!`);
  
  let presenceObj = new Presence(client, {
    user : client.user, 
    game: { name: '!help for commands' , type: 'LISTENING' }, 
    status: 'active' 
  });
  
  // client.user.Presence(presenceObj)
  //   .then(console.log("We are ready for input, Dave."))
  //   .catch(console.error);
  
  // TODO Presence is broken fix it.
  client.user.Presence = presenceObj;
  
});

// Event listener when a user sends a message in the chat.
client.on("messageCreate", msg => {
  // TODO remove debug code
  console.log(`Message content: ${msg.content}`);
  
  // We check the message content and parse it
  let parsedMessage = parseMessage(msg.content);
  
  if (parsedMessage && parsedMessage.type === "text") {
    msg.channel.send({
      content : parsedMessage.content
    });
  } else if (parsedMessage && parsedMessage.type === "embed") {
    msg.channel.send(runeToMessage(parsedMessage.content));
  } else if (parsedMessage && parsedMessage.type === "runeArray") {
    console.log("Rune Array Detected");
    // Default Rune array
    parsedMessage.content.forEach(runeObj => {
      msg.channel.send(runeToMessage(runeObj));
    });
  } else if (parsedMessage && parsedMessage.type === "allRunesLinks") {
    // All rune array
    console.log("Oh, time for all the rune links.");
    msg.channel.send({
      content : "Here are links to descriptions and proper spelling for all the runes.",
      components : parsedMessage.content.components
    })
  }
  }
);

// Initialize bot by connecting to the server
try {
  client.login(process.env.DISCORD_TOKEN);
} catch(err) {
  throw new Error(err)
}