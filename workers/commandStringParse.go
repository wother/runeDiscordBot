package workers

import (
	"strings"
)

const (
	commandIndicatorCharacter = "!"
	maxVerbLength               = 4
)

var numberStrings = []string{"one", "two", "three", "four", "five"}

// ParseMessage parses a message from Discord and returns a response.
func ParseMessage(message string) map[string]interface{} {
	if !isCommandToBot(message) {
		return nil
	}

	// Slice the command indicator off, as it is no longer needed.
	// Command indicators are the first character in the string.
	stringToParse := message[1:]

	// snag an Array of the whole parsed message, split on spaces.
	// also drop it to lowercase.
	inputMessageArr := strings.Fields(strings.ToLower(stringToParse))

	return parseVerb(inputMessageArr)
}

func isCommandToBot(message string) bool {
	return strings.HasPrefix(strings.TrimSpace(message), commandIndicatorCharacter)
}

func parseVerb(inputStringArr []string) map[string]interface{} {
	verb := inputStringArr[0]
	output := make(map[string]interface{})

	switch {
	case strings.HasPrefix(verb, "help"), strings.HasPrefix(verb, "?"), (verb == "info" && len(inputStringArr) == 1):
		helpString := `The Rune Secrets Bot will draw runes for you from the Elder Futhark.

**Commands are:**

!help for this help text
!cast or !cast one for a single rune casting
!cast three for a three rune casting
!cast five for a five rune casting (careful...)
!info allrunes or names or all or list for a list of all the rune names.
!info [runeName] for information on a specific Rune.`

		output["type"] = "text"
		output["content"] = helpString
	case strings.HasPrefix(verb, "cast"), strings.HasPrefix(verb, "draw"), strings.HasPrefix(verb, "rune"):
		if len(inputStringArr) > 1 {
			numStr := inputStringArr[1]
			if stringInSlice(numStr, numberStrings) {
				num := NumStringToInt(numStr)
				if num == 1 {
					output["type"] = "embed"
					output["content"] = RandomRune(1)
				} else {
					output["type"] = "runeArray"
					output["content"] = RandomRune(num)
				}
			}
		} else {
			output["type"] = "embed"
			output["content"] = RandomRune(1)
		}
	case strings.HasPrefix(verb, "info"):
		if len(inputStringArr) > 1 {
			subject := inputStringArr[1]
			if listCommand(subject) {
				output["type"] = "allRunesLinks"
				output["content"] = RuneInfo("names")
			} else if IsRuneName(subject) {
				output["type"] = "info"
				output["content"] = RuneInfo(subject)
			}
		}
	case strings.HasPrefix(verb, "uptime"):
		uptimeString := "All you need to know is I am online. Fer realsies."
		output["type"] = "text"
		output["content"] = uptimeString
	}

	return output
}

func getRune(inputString string, infoBoolean bool) map[string]interface{} {
	output := make(map[string]interface{})
	if stringInSlice(inputString, numberStrings) {
		if inputString == "one" {
			output["content"] = RandomRune(NumStringToInt(inputString))
			output["type"] = "embed"
		} else {
			output["content"] = RandomRune(NumStringToInt(inputString))
			output["type"] = "runeArray"
		}
	} else if IsRuneName(inputString) && !infoBoolean {
		output["content"] = RuneInfo(inputString)
		output["type"] = "embed"
	} else if IsRuneName(inputString) && infoBoolean {
		// Handle rune info image
	} else if listCommand(inputString) {
		output["content"] = RuneInfo("names")
		output["type"] = "allRunesLinks"
	}
	return output
}

func listCommand(listTestString string) bool {
	switch listTestString {
	case "allrunes", "names", "all", "list":
		return true
	default:
		return false
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}