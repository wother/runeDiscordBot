/**
 * Where we will be directing the flow of our output.
 * 
 * This class will take in the message text, and return the
 * runic object (or objects) requested.
 */

const { randomRune, runeInfo, isRuneName, runeInfoImage } = require("./runes.js");
const StringWorkers = require("./stringWorkers.js");

// Constants
const COMMAND_INDICATOR_CHARACTER = "!";
const NUMBER_STRINGS_ARRAY = ["one", "two", "three", "four", "five"];
const MAX_VERB_LENGTH = 4;

function parseMessage(inputMessageString) {
    if (isCommandToBot(inputMessageString)) {
        // Slice the command indicator off, as it is no longer needed.
        // Command indicators are the first character in the string.
        let stringToParse = inputMessageString.slice(1);
        // snag an Array of the whole parsed message, split on spaces.
        // also drop it to lowercase.
        let inputMessageArr = stringToParse.toLowerCase().split(" ");

        return parseVerb(inputMessageArr);
    }
}

function isCommandToBot(inputString) {
    return inputString.trim().startsWith(COMMAND_INDICATOR_CHARACTER);
}

function parseVerb(inputStringArr) {
    // we trim the input before being passed here, this is a "safe" input.
    let verb = inputStringArr[0];
    let output = {
        "content" : "",
        "type" : ""
    }
    // The verbs are what we tell the bot to do.
    if (verb.startsWith("help") || verb.startsWith("?") || (verb === "info" && !inputStringArr[1])) {
        let helpString = `The Rune Secrets Bot will draw runes for you from the elder Futhark.\n
            Commands are: \n
            **!help** for this help text\n
            **!cast** or **!cast one** for a single rune casting\n
            **!cast three** for a three rune casting\n
            **!cast five** for a five rune casting (careful...)\n
            **!info allrunes** or **names** or **all** or **list** for a list of all the rune names.\n
            **!info [runeName]** for information on a specific Rune.\n
            You can also use "cast" for "rune" or "draw" above.
            `;
        let outputObj = { "content": helpString, "type": "text"};
        return outputObj;
    } else if (verb.startsWith("cast") ||
        verb.startsWith("draw") ||
        verb.startsWith("rune")) {

        if ((verb === "cast" || verb === "draw" || verb === "rune") && !inputStringArr[1]) {
            output = { "content" : randomRune(1), "type": "embed" };
        } else if (verb.length > MAX_VERB_LENGTH || inputStringArr.length > 1) {
            let getRuneNumberString = "";
            
            // Testing for brackets.
            if (StringWorkers.hasBrackets(verb)) {
                verb = StringWorkers.removeBrackets(verb);
            } else if (inputStringArr[1] && StringWorkers.hasBrackets(inputStringArr[1])) {
                inputStringArr[1] = StringWorkers.removeBrackets(inputStringArr[1]);
            }

            // Parsing Number Strings
            if (verb.length > MAX_VERB_LENGTH) {
                getRuneNumberString = verb.substr(MAX_VERB_LENGTH).trim();
            } else  if (inputStringArr[1] && verb.length === MAX_VERB_LENGTH) {
                getRuneNumberString = inputStringArr[1];
            }
            output = getRune(getRuneNumberString, false);
        }
    } else if (verb.startsWith("info")) {
        output = getRune(inputStringArr[1] || verb.substr(4).trim(), true);
    } else if (verb.startsWith("uptime")) {
        let uptimeString = `All **you** need to know is I am online.`;
        output = { "content" : uptimeString, "type": "text"};
    }
    return output;
}

function getRune (inputString, infoBoolean) {
    if (NUMBER_STRINGS_ARRAY.includes(inputString)) {
        if (inputString === "one") {
            return { 
                "content" : randomRune(StringWorkers.numStringToInt(inputString)), 
                "type" : "embed"
            };
        } else {
            return { 
                "content" : randomRune(StringWorkers.numStringToInt(inputString)), 
                "type": "runeArray"
            };
        }
    } else if (StringWorkers.isNumber(inputString)) {
        let inputNumber = Number.parseInt(inputString);
        if (inputNumber === 1) {
            return { 
                "content" : randomRune(inputNumber), 
                "type" : "embed"
            };
        } else {
            return { 
                "content" : randomRune(inputNumber), 
                "type": "runeArray"
            };
        }
    } else if (isRuneName(inputString) && !infoBoolean) {
        return { 
            "content" : runeInfo(inputString), 
            "type": "embed"
        };
    } else if (isRuneName(inputString) && infoBoolean) {
        return { 
            "content" : runeInfoImage(inputString), 
            "type": "embed" 
        };
    } else if (listCommand(inputString)) {
        return { 
            "content" : runeInfo("names"), 
            "type" : "text" 
        };
    }
}

function listCommand (listTestString) {
    let output = false;
    switch (listTestString) {
        case "allrunes":
        case "names":
        case "all":
        case "list":
            output = true;
            break;
        default: output = false;
    }
    return output;
}

module.exports = parseMessage;