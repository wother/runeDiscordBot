/**
 * Where we will be directing the flow of our output.
 * 
 * This class will take in the message text, and return the
 * runic object (or objects) requested.
 */

const { randomRune, runeInfo, isRuneName } = require("./runes.js");

// Constants
const COMMAND_INDICATOR_CHARACTER = "!";
const NUMBER_STRINGS_ARRAY = ["one", "two", "three", "four", "five"];

function parseMessage(inputMessageString) {
    if (isCommandToBot(inputMessageString)) {
        // TODO: remove debug code:
        console.log(`Parsing command ${inputMessageString}`);
        // Slice the command indicator off, as it is no longer needed.
        // Command indicators are the first character in the string.
        let stringToParse = inputMessageString.slice(1);
        // snag an Array of the whole parsed message, split on spaces.
        let inputMessageArr = stringToParse.split(" ");

        return parseVerb(inputMessageArr);
    }
}

function isCommandToBot(inputString) {
    return inputString.trim().startsWith(COMMAND_INDICATOR_CHARACTER);
}

function parseVerb(inputStringArr) {
    // input here comes in trimmed without the "!"
    let verb = inputStringArr[0];
    let output = {
        "content" : "",
        "type" : ""
    }
    // The verbs are what we tell the bot to do.
    if (verb.startsWith("help") || verb.startsWith("?")) {
        let helpString = `The Rune Secrets Bot will draw runes for you from the elder Futhark.\n
            Commands are: \n
            **!help** for this help text\n
            **!cast** or **!cast one** for a single rune casting\n
            **!cast three** for a three rune casting\n
            **!cast five** for a five rune casting (careful...)\n
            **!info allrunes** or **names** or **all** for a list of all the rune names.\n
            **!info [runeName]** for information on a specific Rune.\n
            You can also use "cast" for "rune" or "draw" above.
            `;
        let outputObj = { "content": helpString, "type": "text"};
        return outputObj;
    } else if (verb.startsWith("cast") ||
        verb.startsWith("draw") ||
        verb.startsWith("rune")) {

        if (verb === "cast" || verb === "draw" || verb === "rune") {
            output = { "content" : randomRune(1), "type": "embed" };
        } else if (verb.length > 4) {
            let subString = verb.substr(4).trim();
            output = getRune(subString);
        }
    } else if (verb.startsWith("info")) {
        output = getRune(inputStringArr[1] || verb.substr(4).trim());;
    } else if (verb.startsWith("uptime")) {
        let uptimeString = `All **you** need to know is I am online.`;
        output = { "content": uptimeString, "type": "text"};
    }
    return output;

    // - !cast*** wildcard processing
    //     - !caststone
    //     - !draw
    //     - !drawstone
    //     - !drawrune
    //     - !rune
    // - !uptime to tell users how long it has been up} inputMessageString 


}

function getRune(inputString){
    if (NUMBER_STRINGS_ARRAY.includes(inputString)) {
        return { "content" : randomRune(numStringToInt(inputString)), "type": "embed"};
    } else if (isRuneName(inputString)){
        return { "content" : runeInfo(inputString), "type": "embed"};
    } else if (inputString === "allrunes" || inputString === "names" || inputString === "all") {
        return { "content" : runeInfo("names"), "type" : "text" };
    }
}

function numStringToInt(numberString) {
    let output = 1;
    // Sorry... I am baking english into the bot.
    switch (numberString) {
        case "one":
            output = 1;
            break;
        case "two":
            output = 2;
            break;
        case "three":
            output = 3;
            break;
        case "four":
            output = 4;
            break;
        case "five":
            output = 5;
            break;
        default:
            break;
    }
    return output;
}

module.exports = parseMessage;