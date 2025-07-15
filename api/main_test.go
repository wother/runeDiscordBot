package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"runeDiscordBot/workers"
	"testing"
)

func TestMain(m *testing.M) {
	// Seed the random number generator to make tests deterministic
	rand.Seed(1)

	// Create a temporary directory for media files
	tmpDir, err := createTestMediaDir()
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	// Set the media path for the workers package to the temp dir
	originalMediaPath := workers.MediaPath
	workers.MediaPath = tmpDir
	defer func() { workers.MediaPath = originalMediaPath }()

	// Run the tests
	os.Exit(m.Run())
}

func TestGetRuneHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/rune", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRuneHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response RuneResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name == "" {
		t.Error("Expected a rune name in the response")
	}

	if response.ImgBase64 == "" {
		t.Error("Expected a base64 encoded image in the response")
	}
}

func TestGetRunesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/runes?num=3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRunesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []RuneResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if len(response) != 3 {
		t.Error("Expected 3 runes in the response")
	}
}

func TestGetInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/info?name=fehu", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getInfoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response RuneResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name != "fehu" {
		t.Error("Expected rune name to be fehu")
	}
}

func TestGetAllRuneNamesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/runeNames", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllRuneNamesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if len(response) != 24 {
		t.Error("Expected 24 rune names in the response")
	}
}

func TestGetStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getStatusHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response["status"] != "online" {
		t.Error("Expected status to be online")
	}
}

