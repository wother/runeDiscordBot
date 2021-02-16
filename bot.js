// Run dotenv
require("dotenv").config();

// Import libraries
const Discord = require("discord.js");
const client = new Discord.Client();

// workers imports
const parseMessage = require("./workers/commandStringParse.js");

// Event listener when bot connects to the server.
client.on("ready", () => {
  console.log(`Logged in as ${client.user.tag}!`);
  client.user
    .setPresence({
      game: { name: "!help for commands", type: "LISTENING" },
      status: "active",
    })
    .then(console.log("We are ready for input, Dave."))
    .catch(console.error);
});

// Event listener when a user sends a message in the chat.
client.on("message", (msg) => {
  // We check the message content and parse it

  // let emoji = msg.match(/(<a?:(.+):\d{17,18}>|:(.+):)/)[1];
  // let testText = toUTF16(msg.content.codePointAt(0));
  // let nativeMessage = msg;
  // console.log(testText);
  // testText.forEach((char) => {
  //   // console.log(char.match(/(<a?:(.+):\d{17,18}>|:(.+):)/));
  //   console.log(toUTF16(char.codePointAt(0)));
  // });

  let parsedMessage = parseMessage(msg.content);
  if (parsedMessage && parsedMessage.type === "text") {
    msg.channel.send(parsedMessage.content);
  } else if (parsedMessage && parsedMessage.type === "embed") {
    let runeEmbed = runeToEmbed(parsedMessage.content, msg);
    msg.channel.send({"embed" : runeEmbed });
  } else if (parsedMessage && parsedMessage.type === "runeArray") {
    parsedMessage.content.forEach((runeObj) => {
      msg.channel.send({"embed": runeToEmbed(runeObj, msg)});
    });
  }
});

function runeToEmbed(runeObject, inputMessage) {
  const embed = new Discord.MessageEmbed(inputMessage, {
    title: runeObject.name,
    url: runeObject.descURL,
    image: {
      url: runeObject.imgURL,
    },
  });

  return embed;
}

// Test code, if werks, move to StringWorker
function toUTF16(codePoint) {
  var TEN_BITS = parseInt("1111111111", 2);
  function u(codeUnit) {
    return "\\u" + codeUnit.toString(16).toUpperCase();
  }

  if (codePoint <= 0xffff) {
    return u(codePoint);
  }
  codePoint -= 0x10000;

  // Shift right to get to most significant 10 bits
  var leadSurrogate = 0xd800 + (codePoint >> 10);

  // Mask to get least significant 10 bits
  var tailSurrogate = 0xdc00 + (codePoint & TEN_BITS);

  return u(leadSurrogate) + u(tailSurrogate);
}

// Initialize bot by connecting to the server
try {
  client.login(process.env.DISCORD_TOKEN);
} catch (err) {
  throw new Error(err);
}
