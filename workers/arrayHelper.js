// https://stackoverflow.com/questions/37644121/splitting-an-array-into-groups
function chunkArrayinGroupsWithCB(arr, size, cb) {
    let result = [];
    let pos = 0;
    while (pos < arr.length) {
        let item = arr.slice(pos, pos + size)
        if (item) {
            result.push(cb(item));
        }
        pos += size;
    }
    return result;
}

module.exports = chunkArrayinGroupsWithCB;