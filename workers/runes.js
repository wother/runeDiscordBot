/**
 * The place where the Futhark lives. Rune JSON and links to art are here.
 */
// const rand = require('./randomizer.js');
const randomFromArray = require('./randomizer.js')

function randomRune(inputNumber) {
    // TODO parse the number to how many runes to "cast"
    let output = "";

    switch (inputNumber) {
        case 1:
            output = genLink(randomFromArray(futharkArray));
            break;
        case 3:
            output = numUniqueRunes(3);
            break;
        case 5:
            output = numUniqueRunes(5);
            break;
        default :
            output = genLink(randomFromArray(futharkArray));
    }

    return output;
}

const futharkArray = [
    "algiz",
    "ansuz",
    "berkano",
    "dagaz",
    "ehwaz",
    "fehu",
    "gebo",
    "hagalaz",
    "ihwaz",
    "inguz",
    "isa",
    "jera",
    "kenaz",
    "laguz",
    "mannaz",
    "nauthiz",
    "othala",
    "perthro",
    "raidho",
    "sowilo",
    "thurisaz",
    "tiwaz",
    "uruz",
    "wunjo"
];

function genLink(runeName) {
    return `https://runesecrets.com/img/${runeName}-100x100.gif`;
}

function genInfoLink (runeName) {
    return `https://runesecrets.com/rune-meanings/${runeName}`;
}

function numUniqueRunes(inputNumber) {
    let output = "";
    let futharkCopy = [...futharkArray];
    let randoRunes = futharkCopy.sort(()=>{ return .5 - Math.random()})
        .slice(0 , inputNumber);
    randoRunes.forEach(runeName => {
        output += (genLink(runeName) + " ")
    });
    return output.trim();
}

module.exports = randomRune;