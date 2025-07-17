package workers

import (
	"fmt"
	"math/rand"
	"os"
)

var futharkNames = []string{
	"algiz",
	"ansuz",
	"berkano",
	"dagaz",
	"ehwaz",
	"fehu",
	"gebo",
	"hagalaz",
	"ihwaz",
	"inguz",
	"isa",
	"jera",
	"kenaz",
	"laguz",
	"mannaz",
	"nauthiz",
	"othala",
	"perthro",
	"raidho",
	"sowilo",
	"thurisaz",
	"tiwaz",
	"uruz",
	"wunjo",
}

// MediaPath is the base path for media files.
var MediaPath = "media"

// Rune holds information about a single rune.
type Rune struct {
	Name    string
	ImgFile string
	DescURL string
}

// GetFutharkArray returns the futharkNames array.
func GetFutharkArray() []string {
	return futharkNames
}

// RandomRune returns a specified number of random runes.
func RandomRune(number int, r *rand.Rand) interface{} {
	switch number {
	case 1:
		return genRuneObject(RandomFromArray(futharkNames, r), r)
	default:
		return numUniqueRunes(number, r)
	}
}

// RuneInfo returns information about a specific rune.
func RuneInfo(runeName string) interface{} {
	if IsRuneName(runeName) {
		return genRuneObject(runeName, nil) // Pass nil for rand.Rand as it's not needed for info
	} else if runeName == "names" {
		return futharkNames
	}
	return nil
}

// IsRuneName checks if a string is a valid rune name.
func IsRuneName(s string) bool {
	for _, r := range futharkNames {
		if r == s {
			return true
		}
	}
	return false
}

func genRuneObject(name string, r *rand.Rand) Rune {
	return Rune{
		Name:    name,
		ImgFile: genImgPath(name, r),
		DescURL: "", // No longer needed
	}
}

func genImgPath(runeName string, r *rand.Rand) string {
	// 5% chance for rare "dreams" variant
	if r != nil && r.Float32() < 0.05 {
		// Try dreams directory first - let fileExists handle the name variations
		dreamsPath := fmt.Sprintf("%s/solo/dreams/%s.png", MediaPath, runeName)
		if fileExists(dreamsPath) {
			fmt.Printf("🌙 Rare dreams variant selected for %s\n", runeName)
			return dreamsPath
		}
	}

	// Default to regular image
	return fmt.Sprintf("%s/solo/%s.png", MediaPath, runeName)
}

func genSmallImgPath(runeName string) string {
	// Try the exact name first
	exactPath := fmt.Sprintf("%s/solo/resized/%s (Small).png", MediaPath, runeName)
	if fileExists(exactPath) {
		return exactPath
	}

	// Try lowercase variant
	lowerPath := fmt.Sprintf("%s/solo/resized/%s (small).png", MediaPath, runeName)
	if fileExists(lowerPath) {
		return lowerPath
	}

	// Fallback to exact name (will fail gracefully if doesn't exist)
	return exactPath
}

func numUniqueRunes(number int, r *rand.Rand) []Rune {
	output := make([]Rune, 0)
	futharkArrayCopy := make([]string, len(futharkNames))
	copy(futharkArrayCopy, futharkNames)

	r.Shuffle(len(futharkArrayCopy), func(i, j int) {
		futharkArrayCopy[i], futharkArrayCopy[j] = futharkArrayCopy[j], futharkArrayCopy[i]
	})

	randoRunes := futharkArrayCopy[:number]

	for _, runeName := range randoRunes {
		output = append(output, genRuneObject(runeName, r))
	}

	return output
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
