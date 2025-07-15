package workers

import (
	"testing"
)

func TestGetFutharkArray(t *testing.T) {
	if len(GetFutharkArray()) != 24 {
		t.Error("Expected 24 runes in futhark array")
	}
}

func TestRandomRune(t *testing.T) {
	// Test for a single rune
	runeObj := RandomRune(1)
	if _, ok := runeObj.(Rune); !ok {
		t.Error("Expected a single Rune object")
	}

	// Test for multiple runes
	runeArray := RandomRune(3).([]Rune)
	if len(runeArray) != 3 {
		t.Error("Expected 3 runes in the array")
	}
}

func TestRuneInfo(t *testing.T) {
	// Test for a valid rune name
	infoObj := RuneInfo("fehu")
	if _, ok := infoObj.(Rune); !ok {
		t.Error("Expected a Rune object for a valid rune name")
	}

	// Test for an invalid rune name
	infoObj = RuneInfo("not-a-rune")
	if infoObj != nil {
		t.Error("Expected nil for an invalid rune name")
	}

	// Test for "names"
	infoObj = RuneInfo("names")
	if _, ok := infoObj.([]string); !ok {
		t.Error("Expected a slice of strings for \"names\"")
	}
}

func TestIsRuneName(t *testing.T) {
	// Test for a valid rune name
	if !IsRuneName("fehu") {
		t.Error("Expected true for a valid rune name")
	}

	// Test for an invalid rune name
	if IsRuneName("not-a-rune") {
		t.Error("Expected false for an invalid rune name")
	}
}
