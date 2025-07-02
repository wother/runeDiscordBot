package workers

import (
	"fmt"
	"math/rand"
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

// Rune holds information about a single rune.
type Rune struct {
	Name    string
	ImgURL  string
	DescURL string
}

// GetFutharkArray returns the futharkNames array.
func GetFutharkArray() []string {
	return futharkNames
}

// RandomRune returns a specified number of random runes.
func RandomRune(number int) interface{} {
	switch number {
	case 1:
		return genRuneObject(RandomFromArray(futharkNames))
	default:
		return numUniqueRunes(number)
	}
}

// RuneInfo returns information about a specific rune.
func RuneInfo(runeName string) interface{} {
	if IsRuneName(runeName) {
		return genRuneObject(runeName)
	} else if runeName == "names" {
		return AllRunesLinks(futharkNames)
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

func genRuneObject(name string) Rune {
	return Rune{
		Name:    name,
		ImgURL:  genImgLink(name),
		DescURL: genInfoLink(name),
	}
}

func genImgLink(runeName string) string {
	return fmt.Sprintf("/img/%s-100x100.gif", runeName)
}

func genInfoLink(runeName string) string {
	return fmt.Sprintf("/rune-meanings/%s", runeName)
}

func numUniqueRunes(number int) []Rune {
	output := make([]Rune, 0)
	futharkArrayCopy := make([]string, len(futharkNames))
	copy(futharkArrayCopy, futharkNames)

	rand.Shuffle(len(futharkArrayCopy), func(i, j int) {
		futharkArrayCopy[i], futharkArrayCopy[j] = futharkArrayCopy[j], futharkArrayCopy[i]
	})

	randoRunes := futharkArrayCopy[:number]

	for _, runeName := range randoRunes {
		output = append(output, genRuneObject(runeName))
	}

	return output
}
