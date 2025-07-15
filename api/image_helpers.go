package api

import (
	"encoding/base64"
	"io/ioutil"
)

// encodeImageToBase64 reads an image file and returns its Base64-encoded representation.
func encodeImageToBase64(filePath string) (string, error) {
	imgBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imgBytes), nil
}
