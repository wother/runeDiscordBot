const { MessageEmbed, MessageButton, MessageActionRow } = require('discord.js');


function runeToMessage (inputRuneObject) {
  let output = {
    ephemeral: true,
    embeds : [runeToEmbed(inputRuneObject)],
    components : [componentRowWith(runeLinkButton(inputRuneObject.name, inputRuneObject.descURL))]
  }
  return output;
}


function runeToEmbed(runeObject) {
  const embed = new MessageEmbed({
    title : runeObject.name,
    url : runeObject.descURL,
    image : {
      "url": runeObject.imgURL
    },
    description : `You drew ${runeObject.name}`
  });

  return embed;
}

function componentRowWith(inputButton) {
  let outputRow = new MessageActionRow()
    .addComponents(inputButton);
  return outputRow;
}

function runeLinkButton (inputTitle, URLinput) {
  let outputButton = new MessageButton()
    .setLabel(`Go to ${inputTitle} info page.`)
    .setURL(URLinput)
    .setStyle('LINK');
  return outputButton;
}

exports.runeToMessage = runeToMessage;