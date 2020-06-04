/**
 * Where we manipulate strings beyond regular JS Methods.
 * Additionally, we wrap a few String methods and calls for
 * parsing our commands and syntax.
 */

const StringWorkers = {
    hasBrackets : (testString) => {
        let regExp = /([\[\]])\w*/g;
        return regExp.test(testString);
    },
    removeBrackets : (inputTestString) => {
        let output = "";
        let regExp = /([\[\]])\w*/g;
        // if (hasBrackets(inputTestString)) {
            let strArray = inputTestString.split("");
            strArray.forEach(stringChar => {
                if (!regExp.test(stringChar)){
                    output += stringChar;
                }
            })
        // } else {
        //     output = inputTestString;
        // }
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
    }
}

module.exports = StringWorkers;