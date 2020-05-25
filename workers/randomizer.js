/** Where we wrap the Math object to randomize some 
 * values. 
 */

let randomizer = {
    /**
     * This method will return one value at random from the input Object.
     * 
     * @param {Object} inputObject the object you want a random attribute from.
     */
    randomAttribute = function (inputObject) {
        const keysArray = [...Object.keys(inputObject)];
        const keysLength = keysArray.length;
        return inputObject[keysArray[randomIntBetween(0, keysLength - 1)]];
    }
};

function randomIntBetween (min, max) {
    return Math.floor(Math.random() * max + min);
}

module.exports = randomizer;