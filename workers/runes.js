/**
 * The place where the Futhark lives. Rune JSON and links to art are here.
 */
const randomFromArray = require('./randomizer.js')

function randomRune(inputNumber) {
    let output;

    // We generate an array of rune objects with name, hyperlink, and image.

    switch (inputNumber) {
        case 1:
            output = genRuneObject(randomFromArray(futharkArray));
            break;
        case 2:
        case 3:
        case 4:
        case 5:
            output = numUniqueRunes(inputNumber);
            break;
    }

    return output;
}

function runeInfo (inputRuneName) {
    // we want to ensure that the rune in question exists in the Futhark(array).
    // Otherwise we drop through returning nothing.
    // TODO: meaningful error text if falure to find rune name.
    if (futharkArray.includes(inputRuneName)){
        return genRuneObject(inputRuneName);
    } else if(inputRuneName === "names") {
        let formattedRuneString = "";
        futharkArray.forEach(runeName =>{
            formattedRuneString += `${runeName} \n`;
        })
        return formattedRuneString;
    }
}

function isRuneName(inputString) {
    return futharkArray.includes(inputString);
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

function genRuneObject (nameInput) {
    let output = {
        "name" : nameInput,
        "imgURL" : genImgLink(nameInput),
        "descURL" : genInfoLink(nameInput)
    }
    return output;
}

function genImgLink(runeName) {
    return `https://runesecrets.com/img/${runeName}-100x100.gif`;
}

function genInfoLink (runeName) {
    return `https://runesecrets.com/rune-meanings/${runeName}`;
}

function numUniqueRunes(inputNumber) {
    let output = [];
    let futharkCopy = [...futharkArray];
    let randoRunes = futharkCopy.sort(()=>{ return .5 - Math.random()})
        .slice(0 , inputNumber);
    
    randoRunes.forEach(runeName => {
        output.push(genRuneObject(runeName));
    });
    return output;
}

module.exports = {randomRune, runeInfo, isRuneName};