package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"runeDiscordBot/workers"
)

// SetupAPI initializes and starts the API server.
func SetupAPI() {
	http.HandleFunc("/api/rune", getRuneHandler)
	http.HandleFunc("/api/runes", getRunesHandler)
	http.HandleFunc("/api/info", getInfoHandler)
	http.HandleFunc("/api/runeNames", getAllRuneNamesHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	fmt.Printf("The API server is up, Dave. live on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getRuneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	runeObj := workers.RandomRune(1)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runeObj)
}

func getRunesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	numStr := r.URL.Query().Get("num")
	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		http.Error(w, "Invalid number of runes", http.StatusBadRequest)
		return
	}

	runeArray := workers.RandomRune(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runeArray)
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	runeName := r.URL.Query().Get("name")
	if runeName == "" {
		http.Error(w, "Rune name is required", http.StatusBadRequest)
		return
	}

	infoObj := workers.RuneInfo(runeName)
	if infoObj == nil {
		http.Error(w, "Rune not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoObj)
}

func getAllRuneNamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	runeNames := workers.GetFutharkArray()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runeNames)
}
