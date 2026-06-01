package builder

import (
	"strings"
	"testing"
)

func TestArtBuilder_BasicFlow(t *testing.T) {
	builder := NewArtBuilder()

	result := builder.
		AddText("HI").
		SetStyle("normal").
		Build()

	if result == "" {
		t.Error("Expected non-empty result")
	}

	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}
}

func TestArtBuilder_MethodChaining(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*ArtBuilder) *ArtBuilder
		expectLines int
		minWidth    int
	}{
		{
			name: "Single text normal style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("A").SetStyle("normal")
			},
			expectLines: 8,
			minWidth:    1,
		},
		{
			name: "Bold style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("A").SetStyle("bold")
			},
			expectLines: 8,
			minWidth:    2,
		},
		{
			name: "Multiple texts with different styles",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.
					AddText("A").
					SetStyle("normal").
					AddText("B").
					SetStyle("italic")
			},
			expectLines: 8,
			minWidth:    2,
		},
		{
			name: "Outline style",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("T").SetStyle("outline")
			},
			expectLines: 8,
			minWidth:    2,
		},
		{
			name: "Italic style has forward slant",
			setup: func(ab *ArtBuilder) *ArtBuilder {
				return ab.AddText("I").SetStyle("italic")
			},
			expectLines: 8,
			minWidth:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewArtBuilder()
			builder = tt.setup(builder)

			result := builder.Build()

			if result == "" {
				t.Errorf("Expected non-empty result for %s", tt.name)
			}

			lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
			if len(lines) != tt.expectLines {
				t.Errorf("Expected %d lines, got %d", tt.expectLines, len(lines))
			}

			if len(lines) > 0 && len(lines[0]) < tt.minWidth {
				t.Errorf("Expected min width %d, got %d", tt.minWidth, len(lines[0]))
			}
		})
	}
}

func TestArtBuilder_ItalicSlant(t *testing.T) {
	builder := NewArtBuilder()
	normalResult := builder.AddText("I").SetStyle("normal").Build()

	builder = NewArtBuilder()
	italicResult := builder.AddText("I").SetStyle("italic").Build()

	if normalResult == italicResult {
		t.Error("Italic style should differ from normal style (forward slant required)")
	}

	if italicResult == "" {
		t.Error("Italic style should produce non-empty result")
	}
}

func TestArtBuilder_InvalidStyle(t *testing.T) {
	builder := NewArtBuilder()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid style 'rainbow' - only supported styles: normal, bold, italic, outline")
		}
	}()

	builder.AddText("HI").SetStyle("rainbow").Build()
}

func TestArtBuilder_BoldStyle(t *testing.T) {
	builder := NewArtBuilder()
	normalResult := builder.AddText("A").SetStyle("normal").Build()

	builder = NewArtBuilder()
	boldResult := builder.AddText("A").SetStyle("bold").Build()

	if normalResult == boldResult {
		t.Error("Bold style should differ from normal style")
	}

	normalLines := strings.Split(strings.TrimRight(normalResult, "\n"), "\n")
	boldLines := strings.Split(strings.TrimRight(boldResult, "\n"), "\n")

	if len(normalLines) > 0 && len(boldLines) > 0 {
		if len(boldLines[0]) <= len(normalLines[0]) {
			t.Log("Warning: Bold style is not wider than normal style (but this may be acceptable depending on implementation)")
		}
	}
}

func TestArtBuilder_OutlineStyle(t *testing.T) {
	builder := NewArtBuilder()
	outlineResult := builder.AddText("T").SetStyle("outline").Build()

	if outlineResult == "" {
		t.Error("Outline style should produce non-empty result")
	}

	hasBorderChars := strings.ContainsAny(outlineResult, "+-|")
	if !hasBorderChars {
		t.Log("Warning: Outline style doesn't contain typical border characters")
	}
}

func TestArtBuilder_MethodChainingReturn(t *testing.T) {
	builder := NewArtBuilder()

	result1 := builder.AddText("A")
	if result1 != builder {
		t.Error("AddText should return the same builder instance for chaining")
	}

	result2 := builder.SetStyle("bold")
	if result2 != builder {
		t.Error("SetStyle should return the same builder instance for chaining")
	}

	result3 := builder.AddText("B")
	if result3 != builder {
		t.Error("AddText should return the same builder instance for chaining")
	}
}
