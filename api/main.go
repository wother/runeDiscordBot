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
	http.HandleFunc("/api/", getDocsHandler)
	http.HandleFunc("/api/docs", getDocsHandler)
	http.HandleFunc("/api/rune", getRuneHandler)
	http.HandleFunc("/api/runes", getRunesHandler)
	http.HandleFunc("/api/info", getInfoHandler)
	http.HandleFunc("/api/runeNames", getAllRuneNamesHandler)
	http.HandleFunc("/api/status", getStatusHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	fmt.Printf("The API server is up, Dave. live on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDocsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	docs := `

The Rune Secrets Bot API allows you to draw runes from the Elder Futhark.

**Endpoints:**

* **GET /api/status**: Returns the API's status.
* **GET /api/rune**: Returns a single random rune with an embedded image.
* **GET /api/runes?num=[number]**: Returns a specified number of random runes with embedded images.
* **GET /api/info?name=[runeName]**: Returns information on a specific Rune, with an embedded image.
* **GET /api/runeNames**: Returns a list of all the rune names.
`
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, docs)
}

func getRuneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	runeObj := workers.RandomRune(1).(workers.Rune)
	imgBase64, err := encodeImageToBase64(runeObj.ImgFile)
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}

	response := RuneResponse{
		Name:      runeObj.Name,
		ImgBase64: imgBase64,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	runeArray := workers.RandomRune(num).([]workers.Rune)
	var response []RuneResponse

	for _, runeObj := range runeArray {
		imgBase64, err := encodeImageToBase64(runeObj.ImgFile)
		if err != nil {
			http.Error(w, "Failed to encode image", http.StatusInternalServerError)
			return
		}
		response = append(response, RuneResponse{
			Name:      runeObj.Name,
			ImgBase64: imgBase64,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	runeObj := infoObj.(workers.Rune)
	imgBase64, err := encodeImageToBase64(runeObj.ImgFile)
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}

	response := RuneResponse{
		Name:      runeObj.Name,
		ImgBase64: imgBase64,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "I'm a teapot, not a Coffee maker, cannot brew coffee in a teapot.", http.StatusTeapot)
		return
	}

	status := map[string]string{"status": "online"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
