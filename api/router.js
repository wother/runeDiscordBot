import {
    getRune, getRunes, getInfo, getAllRuneNames
} from "./controllers.js";

export async function router(req, res) {
    // Parse out the arguments and seperate from the base URL encoded in the request
    let baseURL, args;
    if (req) {
        let urlArray = req.url.split("?");
        baseURL = urlArray[0];
        args = urlArray[1];
    }
    // Switch for calling controller methods
    switch (baseURL) {
        case '/api/rune':
            // Handle rune endpoint
            let runeObj = getRune();
            writeStringifyEndRes(res, runeObj)
        case '/api/runes':
            // Handle multi rune endpoint
            let runeArray = getRunes(Number.parseInt(args));
            writeStringifyEndRes(res, runeArray);
        case '/api/info':
            // Handle information request
            let infoObj = getInfo(args);
            writeStringifyEndRes(res, infoObj);
        case '/api/runeNames':
            let runeNameArray = getAllRuneNames();
            writeStringifyEndRes(res, runeNameArray);
        default:
            // Handle fallthrough
    }
}

function writeStringifyEndRes(responseInput, inputJson) {
    responseInput.write(JSON.stringify(inputJson));
    responseInput.end();
}
