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
            console.log(runeObj);
            writeStringifyEndRes(res, runeObj);
            break;
        case '/api/runes':
            // Handle multi rune endpoint
            let runeArray = getRunes(Number.parseInt(args));
            console.log(runeArray)
            writeStringifyEndRes(res, runeArray);
            break;
        case '/api/info':
            // Handle information request
            let infoObj = getInfo(args);
            console.log(infoObj)
            writeStringifyEndRes(res, infoObj);
            break;
        case '/api/runeNames':
            let runeNameArray = getAllRuneNames();
            console.log(runeNameArray)
            writeStringifyEndRes(res, runeNameArray);
            break;
        default:
            // Handle fallthrough
    }
};

function writeStringifyEndRes(responseInput, inputJson) {
    responseInput.write(JSON.stringify(inputJson));    
    responseInput.end();
};
