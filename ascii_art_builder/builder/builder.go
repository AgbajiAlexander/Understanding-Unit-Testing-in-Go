package builder

import (
	"strings"
)

type segment struct {
	text  string
	style string
}

type ArtBuilder struct {
	segments []segment
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.segments = append(ab.segments, segment{
		text:  text,
		style: "normal",
	})
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	switch style {
	case "normal", "bold", "italic", "outline":
	default:
		panic("unsupported style")
	}

	if len(ab.segments) > 0 {
		ab.segments[len(ab.segments)-1].style = style
	}

	return ab
}

func (ab *ArtBuilder) Build() string {
	lines := make([]string, 8)

	for _, seg := range ab.segments {
		rendered := render(seg.text, seg.style)

		for i := 0; i < 8; i++ {
			lines[i] += rendered[i]
		}
	}

	return strings.Join(lines, "\n") + "\n"
}

func render(text, style string) []string {
	lines := make([]string, 8)

	for row := 0; row < 8; row++ {
		var line strings.Builder

		for _, ch := range text {
			line.WriteString(styleChar(ch, style, row))
		}

		lines[row] = line.String()
	}

	return lines
}

func styleChar(ch rune, style string, row int) string {
	base := string(ch)

	switch style {
	case "normal":
		return base

	case "bold":
		return base + base

	case "italic":
		// Forward slant using leading spaces
		spaces := strings.Repeat(" ", 7-row)
		return spaces + base

	case "outline":
		return "|" + base + "|"

	default:
		panic("unsupported style")
	}
}
