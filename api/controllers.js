import {
    randomRune
} from "../workers/runes.js";


export function getRune() {
    // return rune object

    return randomRune(1);
}

export function getRunes(numberOfRunes) {
    // return JSON array of rune objects
}

export function getInfo(inputRuneName) {
    // return the rune info JSON
}