/**
 * Where we manipulate strings beyond regular JS Methods.
 * Additionally, we wrap a few String methods and calls for
 * parsing our commands and syntax.
 */

// The Regular Expression to test if a string has brackets in it.
const BRACKET_REGEX = /([\[\]])\w*/g;
const NUMBER_REGEX = /[\d]/g;
const COLON_REGEX = /([\:])\w*/g;

const StringWorkers = {
    hasBrackets : (testString) => {
        return BRACKET_REGEX.test(testString);
    },
    removeBrackets : (inputTestString) => {
        let output = "";
        let strArray = inputTestString.split("");
        strArray.forEach(stringChar => {
            if (!(stringChar.includes("[") || stringChar.includes("]"))){
                output += stringChar;
            }
        });
        return output;
    },
    hasColon : (testString) => {
        return COLON_REGEX.test(testString);
    },
    removeColons : (inputTestString) => {
        let output = "";
        let strArray = inputTestString.split("");
        strArray.forEach(stringChar => {
            if (!stringChar.includes(":")){
                output += stringChar;
            }
        });
        return output;
    },
    numStringToInt : (numberString) => {
        let output = 1;
        // English enumeration.
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
    },
    capitalize : (inputString) => {
        let snippedStr = inputString.substr(1).toLowerCase(); 
        let firstLetter = inputString.charAt(0).toString().toUpperCase();
        return firstLetter + snippedStr;
    },
    isNumber : (inputNumberString) => {
        return NUMBER_REGEX.test(inputNumberString); 
    }
}

module.exports = StringWorkers;