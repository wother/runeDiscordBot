export function router(req, res) {
    switch (req.url) {
        case '/api/rune':
            // Handle rune endpoint
        case '/api/runes':
            // Handle multi rune endpoint
        case '/api/info':
            // Handle information request
        default:
            // Handle fallthrough
    }
}