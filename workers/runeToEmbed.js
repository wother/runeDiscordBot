const { MessageEmbed, MessageButton, MessageActionRow } = require('discord.js');
const chunkArrayinGroupsWithCB = require('./arrayHelper.js');

function runeToMessage (inputRuneObject) {
  let output = {
    ephemeral: true,
    embeds : [runeToEmbed(inputRuneObject)],
    components : [componentRowWith(runeLinkButton(inputRuneObject.name, inputRuneObject.descURL))]
  }
  return output;
}

function allRunesLinks(futharkArray) {
  // hard limit of five by five here
  // five Action Rows, with Five Buttons each
  let output = {
    ephemeral: true,
    components : []
  };
  let runeButtonMatrix = chunkArrayinGroupsWithCB(futharkArray, 5, runeNameButton);
  runeButtonMatrix.forEach(buttonArray => {
    let messageRow = new MessageActionRow();
    buttonArray.forEach(runeButton => {
      messageRow.addComponents(runeButton);
    });
    output.components.push(messageRow);
  });
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

function runeNameButton(name) {
  let outputButton = new MessageButton()
  .setLabel(name)
  .setURL(genInfoLink(name))
  .setStyle('LINK');
return outputButton;
}

exports = { runeToMessage, allRunesLinks };