import { MessageEmbed, MessageButton, MessageActionRow } from 'discord.js';
import {chunkArrayinGroupsWithCB} from './arrayHelper.js';
import { genInfoLink } from './runes.js';

export function runeToMessage (inputRuneObject) {
  let output = {
    ephemeral: true,
    embeds : [runeToEmbed(inputRuneObject)],
    components : [componentRowWith(runeLinkButton(inputRuneObject.name, inputRuneObject.descURL))]
  }
  return output;
}

export function allRunesLinks(futharkArray) {
  // hard limit of five by five here
  // five Action Rows, with Five Buttons each
  let output = {
    ephemeral: true,
    components : []
  };
  let runeNameMatrix = chunkArrayinGroupsWithCB(futharkArray, 5);
  runeNameMatrix.forEach(nameArray => {
    let messageRow = new MessageActionRow();
    nameArray.forEach(runeName => {
      messageRow.addComponents(runeNameButton(runeName));
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
  .setLabel(name.toString())
  .setURL(genInfoLink(name))
  .setStyle('LINK');
return outputButton;
}