package converter

import "strings"

// Each digit is 5 chars wide x 5 lines tall
var digits = map[rune][5]string{
	'0': {
		" ___ ",
		"|   |",
		"|   |",
		"|   |",
		"|___|",
	},
	'1': {
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
		"  |  ",
	},
	'2': {
		" ___ ",
		"    |",
		" ___|",
		"|    ",
		"|___|",
	},
	'3': {
		" ___ ",
		"    |",
		" ___|",
		"    |",
		"|___|",
	},
	'4': {
		"     ",
		"|   |",
		"|___|",
		"    |",
		"    |",
	},
	'5': {
		" ___ ",
		"|    ",
		"|___ ",
		"    |",
		"|___|",
	},
	'6': {
		" ___ ",
		"|    ",
		"|___ ",
		"|   |",
		"|___|",
	},
	'7': {
		" ___ ",
		"    |",
		"    |",
		"    |",
		"    |",
	},
	'8': {
		" ___ ",
		"|   |",
		"|___|",
		"|   |",
		"|___|",
	},
	'9': {
		" ___ ",
		"|   |",
		"|___|",
		"    |",
		"|___|",
	},
}

// StringToArt converts a string of digits into ASCII art.
// Supports digits 0-9, multiple lines separated by \n.
// Returns empty string for invalid input.
func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	lines := strings.Split(input, "\n")
	var result strings.Builder

	for lineIdx, line := range lines {
		if line == "" {
			return ""
		}

		// Validate: only digits allowed
		for _, ch := range line {
			if ch < '0' || ch > '9' {
				return ""
			}
		}

		// Build 5 rows for this line of digits
		for row := 0; row < 5; row++ {
			for _, ch := range line {
				result.WriteString(digits[ch][row])
			}
			result.WriteString("\n")
		}

		_ = lineIdx
	}

	return result.String()
}
