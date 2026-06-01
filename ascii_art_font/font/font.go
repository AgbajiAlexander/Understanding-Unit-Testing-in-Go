package font

import "strings"

// GenerateFont creates all printable ASCII characters (32-126)
// as 8x8 algorithmically-generated glyphs.
func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for r := rune(32); r <= rune(126); r++ {
		font[r] = generateGlyph(r)
	}

	return font
}

func generateGlyph(r rune) []string {
	// Space must be entirely blank
	if r == ' ' {
		return blankGlyph()
	}

	grid := make([][]rune, 8)
	for i := range grid {
		grid[i] = []rune("........")
	}

	code := int(r)

	// ----------------------------------
	// Unique fingerprint from ASCII code
	// ----------------------------------

	// Primary diagonal pattern
	for row := 0; row < 8; row++ {
		col := (row + code) % 8
		grid[row][col] = '*'
	}

	// Bit-driven pattern
	for bit := 0; bit < 8; bit++ {
		if (code>>bit)&1 == 1 {
			col := (code + bit) % 8
			grid[bit][col] = '*'
		}
	}

	// ----------------------------------
	// Character categories
	// ----------------------------------

	if isLetter(r) {
		if isVowel(r) {
			addBorder(grid)
		} else {
			addLeftStem(grid)
		}
	}

	if isAscender(r) {
		addAscender(grid)
	}

	if isDescender(r) {
		addDescender(grid)
	}

	if isDigit(r) {
		addDigitMark(grid)
	}

	if isPunctuation(r) {
		addCross(grid)
	}

	lines := make([]string, 8)
	for i := range grid {
		lines[i] = string(grid[i])
	}

	return lines
}

func blankGlyph() []string {
	lines := make([]string, 8)

	for i := range lines {
		lines[i] = "        "
	}

	return lines
}

func addBorder(grid [][]rune) {
	for i := 0; i < 8; i++ {
		grid[0][i] = '*'
		grid[7][i] = '*'
		grid[i][0] = '*'
		grid[i][7] = '*'
	}
}

func addLeftStem(grid [][]rune) {
	for i := 0; i < 8; i++ {
		grid[i][1] = '*'
	}
}

func addAscender(grid [][]rune) {
	for row := 0; row < 3; row++ {
		grid[row][2] = '*'
	}
}

func addDescender(grid [][]rune) {
	for row := 5; row < 8; row++ {
		grid[row][5] = '*'
	}
}

func addDigitMark(grid [][]rune) {
	for i := 0; i < 8; i++ {
		grid[i][i] = '*'
	}
}

func addCross(grid [][]rune) {
	for i := 0; i < 8; i++ {
		grid[3][i] = '*'
		grid[i][3] = '*'
	}
}

func isLetter(r rune) bool {
	return ('A' <= r && r <= 'Z') ||
		('a' <= r && r <= 'z')
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isVowel(r rune) bool {
	return strings.ContainsRune("AEIOUaeiou", r)
}

func isAscender(r rune) bool {
	return strings.ContainsRune("bdfhkltBDFHKLT", r)
}

func isDescender(r rune) bool {
	return strings.ContainsRune("gjpqyGJPQY", r)
}

func isPunctuation(r rune) bool {
	return !isLetter(r) &&
		!isDigit(r) &&
		r != ' '
}
