// https://stackoverflow.com/questions/37644121/splitting-an-array-into-groups
export function chunkArrayinGroupsWithCB(arr, size) {
    let result = [];
    let pos = 0;
    while (pos < arr.length) {
        let item = arr.slice(pos, pos + size)
        if (item) {
            result.push(item);
        }
        pos += size;
    }
    return result;
}