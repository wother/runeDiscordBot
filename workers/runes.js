/**
 * The place where the Futhark lives. Rune JSON and links to art are here.
 */
// const rand = require('./randomizer.js');


function randomRune(inputNumber) {
    return `Rune Number ${inputNumber}`;
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
    return `https://runesecrets.com/img/${runeName}-50x50.gif`;
}

module.exports = randomRune;