import {
    randomRune, runeInfo, getFutharkArray
} from "../workers/runes.js";


export function getRune() {
    // return rune object
    return randomRune(1);
};

export function getRunes(numberOfRunes) {
    // return JSON array of rune objects
    return randomRune(numberOfRunes);
};

export function getInfo(inputRuneName) {
    // return the rune info JSON
    return runeInfo(inputRuneName);
};

export function getAllRuneNames() {
    // Return teh array of Rune Names
    return getFutharkArray();
};