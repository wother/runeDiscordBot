import {} from 'dotenv/config';
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
import { runeToMessage } from "./workers/runeToEmbed.js";
import { parseMessage } from "./workers/commandStringParse.js";

// API Import
import { setupAPI } from "./api/app.js";

// Event listener when bot connects to the server.
client.on("ready", () => {
  console.log(`Logged in as ${client.user.tag}!`);
  
  let presenceObj = new Presence(client, {
    user : client.user, 
    game: { name: '!help for commands' , type: 'LISTENING' }, 
    status: 'active' 
  });
  // TODO Presence is broken, fix it.
  client.user.Presence = presenceObj;
  
});

// Event listener when a user sends a message in the chat.
client.on("messageCreate", msg => {
  // We check the message content and parse it
  let parsedMessage = parseMessage(msg.content);
  
  if (parsedMessage && parsedMessage.type === "text") {
    msg.channel.send({
      content : parsedMessage.content
    });
  } else if (parsedMessage && parsedMessage.type === "embed") {
    msg.channel.send(runeToMessage(parsedMessage.content));
  } else if (parsedMessage && parsedMessage.type === "runeArray") {
    // Default Rune array
    parsedMessage.content.forEach(runeObj => {
      msg.channel.send(runeToMessage(runeObj));
    });
  } else if (parsedMessage && parsedMessage.type === "allRunesLinks") {
    // All rune array
    msg.channel.send({
      content : "Here are links to descriptions and proper spelling for all the runes.",
      components : parsedMessage.content.components
    })
  }
  }
);

// Initialize bot by connecting to the server
try {
  client.login(process.env.TEST_BOT_TOKEN);
  // client.login(config.token);
  setupAPI();
} catch (err) {
  throw new Error(err);
}
