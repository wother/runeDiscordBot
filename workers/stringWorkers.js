/**
 * Where we manipulate strings beyond regular JS Methods.
 * Additionally, we wrap a few String methods and calls for
 * parsing our commands and syntax.
 */

// The Regular Expression to test if a string has brackets in it.
const BRACKET_REGEX = /([\[\]])\w*/g;


export function hasBrackets (testString) {
    return BRACKET_REGEX.test(testString);
}
export function removeBrackets (inputTestString) {
    let output = "";
    let strArray = inputTestString.split("");
    strArray.forEach(stringChar => {
        if (!BRACKET_REGEX.test(stringChar)){
            output += stringChar;
        }
    });
    return output;
}
export function numStringToInt (numberString) {
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
}