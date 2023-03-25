/**
 * This method will return one value at random from the input Array.
 * 
 * @param {Array} inputObject the array you want a random value from.
 */
 export function randomFromArray (arrayInput) {
    return arrayInput[randomIntBetween(0, arrayInput.length - 1)];
};

function randomIntBetween (min, max) {
    return Math.floor(Math.random() * max + min);
}