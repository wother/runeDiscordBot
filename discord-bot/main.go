package main

import (
    "log"
    "net/http"
    "os"

    "discord-bot/internal/discord"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8888"
    }

    // Create a new ServeMux instead of using the default one
    mux := http.NewServeMux()

    // Register your handler on the custom mux
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            discord.HandleRequest(w, r)
        } else {
            w.WriteHeader(http.StatusTeapot)
            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(`{"message": "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot."}`))
        }
    })

    // Pass the custom mux to ListenAndServe
    log.Printf("The API server is up, Dave. Live on port %s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}