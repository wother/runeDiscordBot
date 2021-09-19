const { MessageEmbed, MessageButton } = require('discord.js');

function runeToEmbed(runeObject) {
  const embed = new MessageEmbed({
    title : runeObject.name,
    url : runeObject.descURL,
    image : {
      "url": runeObject.imgURL
    },
    buttons : [runeLinkButton(runeObject.name, runeObject.descURL)],
    description : `You drew ${runeObject.name}`
  });

  return embed;
}

function runeLinkButton (inputTitle, URLinput) {
  let outputButton = new MessageButton(
    {
      label : inputTitle,
      url : URLinput
    }
  );
  return outputButton;
}

exports.runeToEmbed = runeToEmbed;