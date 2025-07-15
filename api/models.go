package api

// RuneResponse represents the structure of a rune API response.
// It includes the rune's name and the Base64-encoded image data.
type RuneResponse struct {
	Name      string `json:"name"`
	ImgBase64 string `json:"img_base64"`
}
