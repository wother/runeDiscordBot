/** Where we wrap the Math object to randomize some 
 * values. 
 */
/**
 * This method will return one value at random from the input Object.
 * 
 * @param {Object} inputObject the object you want a random attribute from.
 */
export function randomAttribute (inputObject) {
    const keysArray = [Object.keys(inputObject)];
    const keysLength = keysArray.length;
    return inputObject[keysArray[randomIntBetween(0, keysLength - 1)]];
 }

 /**
 *  Returns an Integer between the two inputs including the inputs.
 * 
 * @param {Number} min smallest Number you want
 * @param {Number} max largets Number you want
 * @returns {Number}
 */
export function randomIntBetween(min, max) {
     return Math.floor(Math.random() * max + min);
 }

// module.exports = {
//     randomIntBetween,
//     randomAttribute
// }