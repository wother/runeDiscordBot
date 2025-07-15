package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runeDiscordBot/workers"
)

// createTestMediaDir creates a temporary media directory with dummy image files for testing.
func createTestMediaDir() (string, error) {
	tmpDir, err := ioutil.TempDir("", "test-media")
	if err != nil {
		return "", err
	}

	soloDir := filepath.Join(tmpDir, "solo")
	if err := os.MkdirAll(soloDir, 0755); err != nil {
		return "", err
	}

	dreamsDir := filepath.Join(tmpDir, "solo", "dreams")
	if err := os.MkdirAll(dreamsDir, 0755); err != nil {
		return "", err
	}

	// Create dummy image files for all runes
	runeNames := workers.GetFutharkArray()
	for _, name := range runeNames {
		// Create regular image
		dummyFile, err := os.Create(filepath.Join(soloDir, name+".png"))
		if err != nil {
			return "", err
		}
		dummyFile.Write([]byte{0}) // Write a single byte to make the file non-empty
		dummyFile.Close()

		// Create dreams image
		dummyDreamsFile, err := os.Create(filepath.Join(dreamsDir, name+".png"))
		if err != nil {
			return "", err
		}
		dummyDreamsFile.Write([]byte{0}) // Write a single byte to make the file non-empty
		dummyDreamsFile.Close()
	}

	return tmpDir, nil
}
