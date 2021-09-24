/**
 * The tests for our randomizer functions.
 * 
 * We are looking for a spread of numbers with a mean value
 * that has few outliers.
 * 
 * A passing test is one that satisfies this criteria.
 * A failing test is one that has too many outliers.
 * 
 * Our constraints are that the tested functions only need to
 * deal with Twenty-Five (25) whole consecutive integers. 
 * 
 * These integers are held in an array, which we will generate here.
 */

const randomFromArray = require("../workers/randomizer.js");

const MIN_ARRAY_SIZE = 25;
const MAX_ARRAY_SIZE = 250;

const MIN_TEST_INTEGER = 1;
const MAX_TEST_INTEGER = 25;

const MARGIN = 1;

/**
 * Takes an array of integers and determines if their mean (average)
 * is within the set margin of error.
 * 
 * margin of error here is (Min value - Max value)/2 +- 1
 * 
 * @param {Array} inputIntArray array of integers
 */
function testForDistribution (inputIntArray) {

    let sortedArray = inputIntArray.sort();
    let averageOfValues = averageValues(inputIntArray);
    let range = Math.abs(sortedArray[0] - sortedArray[sortedArray.length - 1]);

    let testResults = { "marginResult"  : marginTest(range ),
                        "boundsResult"  : boundsTest(range / 2, averageOfValues),
                        "testArray"     : inputIntArray
        }

    return testResults;
}

function marginTest (inputValue) {
    
    let testValue = Math.abs(inputValue) - MARGIN;

    let output = {"result": "fail", "testValue": testValue};

    if (testValue <= 1){
        output = {"result" : "pass", "testValue" : testValue};
    }
    return output;
}

function boundsTest(val1, val2) {

}

function averageValues(inputArray) {
    let total = 0;
    let length = inputArray.length;
    inputArray.forEach(value => {
        total += value;
    });
    return total / length;
}

function genRandomArray () {
    let outputArray = [];
    
    let randomLengthInt = softRandomNumber(MIN_ARRAY_SIZE, MAX_ARRAY_SIZE);
    
    for (let index = 0; index < randomLengthInt; index++) {
        outputArray.push(softRandomNumber(MIN_TEST_INTEGER, MAX_TEST_INTEGER));
    }

    return outputArray;
};

function softRandomNumber (min, max) {
    return  Math.floor(Math.random() * max + min);
}