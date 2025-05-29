/**
 * The place where the Futhark lives. Rune JSON and links to art are here.
 */
import { randomFromArray } from './randomizer.js';
import { allRunesLinks } from './runeToEmbed.js';

const FUTHARK_NAMES_ARRAY = [
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

export function getFutharkArray() {
    return FUTHARK_NAMES_ARRAY;
}

export function randomRune(inputNumber) {
    let output;

    // We generate an array of rune objects with name, hyperlink, and image.

    switch (inputNumber) {
        case 1:
            output = genRuneObject(randomFromArray(FUTHARK_NAMES_ARRAY));
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

export function runeInfo(inputRuneName) {
    // we want to ensure that the rune in question exists in the Futhark(array).
    // Otherwise we drop through returning nothing.
    
    // TODO: meaningful error text if falure to find rune name.
    
    if (FUTHARK_NAMES_ARRAY.includes(inputRuneName)) {
        return genRuneObject(inputRuneName);
    } else if (inputRuneName === "names") {
        return allRunesLinks(FUTHARK_NAMES_ARRAY);
    }
}

export function isRuneName(inputString) {
    return FUTHARK_NAMES_ARRAY.includes(inputString);
}

function genRuneObject(nameInput) {
    let output = {
        "name": nameInput,
        "imgURL": genImgLink(nameInput),
        "descURL": genInfoLink(nameInput)
    }
    return output;
}

export function runeInfoImage(nameInput) {
    let infoNameString = `${nameInput} information page.`
    let output = {
        "name": infoNameString,
        "imgURL": getInfoImage(nameInput),
        "descURL": genInfoLink(nameInput)
    }
    return output;
}

function genImgLink(runeName) {
    return `/img/${runeName}-100x100.gif`;
}

export function getInfoImage(runeName) {
    return `/img/info-${runeName}.gif`;
}

export function genInfoLink(runeName) {
    return `/rune-meanings/${runeName}`;
}

function numUniqueRunes(inputNumber) {
    let output = [];
    let futharkArrayCopy = [...FUTHARK_NAMES_ARRAY];
    let randoRunes = futharkArrayCopy.sort(() => {
            return .5 - Math.random()
        }).slice(0, inputNumber);

    randoRunes.forEach(runeName => {
        output.push(genRuneObject(runeName));
    });
    return output;
}
