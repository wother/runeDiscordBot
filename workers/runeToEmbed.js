const { MessageEmbed, MessageButton } = require('discord.js');

function runeToEmbed(runeObject, inputMessage) {

  const embed = new MessageEmbed(inputMessage, {
    "title": runeObject.name,
    "url": runeObject.descURL,
    "image": {
      "url": runeObject.imgURL
    },
    "buttons" : [runeLinkButton(runeObject.name, runeObject.descURL)]
  });

  return embed;
}
exports.runeToEmbed = runeToEmbed;

function runeLinkButton (inputTitle, URLinput) {
  let outputButton = new MessageButton();
  outputButton.setLabel(inputTitle);
  outputButton.setURL(URLinput);
  return outputButton;
}