package main

import (
	"fmt"
		"os"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"runeDiscordBot/internal/workers"
	"runeDiscordBot/tui"
)

type appState int

const (
	mainMenu appState = iota
	singleRuneDisplay
	multiRuneDisplay
)

type model struct {
	state            appState
	menuChoices      []string
	cursor           int
	selectedRunes    []workers.Rune
	randSource       rand.Source
	randGenerator    *rand.Rand
	lastDrawnCount   int // New field to store the count of the last drawn runes
}

func initialModel() model {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return model{
		state:         mainMenu,
		menuChoices:   []string{"Draw a single rune", "Draw 3 runes", "Draw 5 runes", "Quit"},
		cursor:        0,
		selectedRunes: []workers.Rune{},
		randSource:    source,
		randGenerator: r,
		lastDrawnCount: 0, // Initialize to 0
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.state == mainMenu {
				if m.cursor > 0 {
					m.cursor--
				}
			}
		case "down", "j":
			if m.state == mainMenu {
				if m.cursor < len(m.menuChoices)-1 {
					m.cursor++
				}
			}

		case "enter":
			if m.state == mainMenu {
				switch m.cursor {
				case 0: // Draw a single rune
					m.selectedRunes = []workers.Rune{workers.RandomRune(1, m.randGenerator).(workers.Rune)}
					m.state = singleRuneDisplay
					m.lastDrawnCount = 1 // Store the count
				case 1: // Draw 3 runes
					m.selectedRunes = workers.RandomRune(3, m.randGenerator).([]workers.Rune)
					m.state = multiRuneDisplay
					m.lastDrawnCount = 3 // Store the count
				case 2: // Draw 5 runes
					m.selectedRunes = workers.RandomRune(5, m.randGenerator).([]workers.Rune)
					m.state = multiRuneDisplay
					m.lastDrawnCount = 5 // Store the count
				case 3: // Quit
					return m, tea.Quit
				}
			} else { // In a display state, pressing enter goes back to menu
				m.state = mainMenu
				m.cursor = 0 // Reset cursor for menu
				m.lastDrawnCount = 0 // Reset count
			}

		case " ": // Spacebar
			if m.state == singleRuneDisplay || m.state == multiRuneDisplay {
				// Redraw the same number of runes
				if m.lastDrawnCount == 1 {
					m.selectedRunes = []workers.Rune{workers.RandomRune(1, m.randGenerator).(workers.Rune)}
				} else if m.lastDrawnCount > 1 {
					m.selectedRunes = workers.RandomRune(m.lastDrawnCount, m.randGenerator).([]workers.Rune)
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	switch m.state {
	case mainMenu:
		s += "Welcome to the Rune Oracle!\n\n"
		s += "What would you like to do?\n\n"
		for i, choice := range m.menuChoices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		s += "\n(Press 'q' or 'ctrl+c' to quit)"

	case singleRuneDisplay:
		runeObj := m.selectedRunes[0]
		runeChar := tui.RuneUnicodeMap[runeObj.Name]
		s += fmt.Sprintf("\n  Your Rune: %s\n\n", runeObj.Name)
		s += fmt.Sprintf("  %s\n\n", runeChar)
		s += "Press Space for another, Enter to return to menu."

	case multiRuneDisplay:
		s += "\nYour Rune Spread:\n\n"
		for _, runeObj := range m.selectedRunes {
			runeChar := tui.RuneUnicodeMap[runeObj.Name]
			s += fmt.Sprintf("  %s: %s\n", runeObj.Name, runeChar)
		}
		s += "\nPress Space for another spread, Enter to return to menu."
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
