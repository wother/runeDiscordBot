import { getRune } from "./controllers.js";

export async function router(req, res) {
    
    
    switch (req.url) {
        case '/api/rune':
            // Handle rune endpoint
            let runeObj = getRune();
            res.write(JSON.stringify(runeObj));
            res.end();
        case '/api/runes':
            // Handle multi rune endpoint
        case '/api/info':
            // Handle information request
        default:
            // Handle fallthrough
    }
}