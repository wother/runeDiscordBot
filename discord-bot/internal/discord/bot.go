package discord

import (
    "encoding/json"
    "net/http"
)

const PORT = "8888"

func init() {
    http.HandleFunc("/", HandleRequest)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Discord bot is running!"})
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusTeapot)
        json.NewEncoder(w).Encode(map[string]string{"message": "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot."})
    }
}

func StartServer() {
    http.ListenAndServe(":"+PORT, nil)
}