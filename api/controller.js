import http from 'http';

const PORT = process.env.PORT || 8888;

async function init(req, res) {
    if (req.url === '/api' && req.method === 'GET') {
        res.writeHead(200, {
            "Content-Type": "application/json"
        });
        res.write("Hello there, Dave.");
        
        res.end()
    } else {
        res.writeHead(418, {
            "Content-Type": "applicaiton/json"
        });
        res.end(JSON.stringify({ message : "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot." }));
    }
}

export function setupAPI() {
    let server = http.createServer(init);
    server.listen(PORT, () => {
        console.log(`The API server is up, Dave. live on port ${PORT}`);
    })
}