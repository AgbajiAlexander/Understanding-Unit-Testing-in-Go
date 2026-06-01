package main

import (
	"fmt"

	"ascii_art_builder/builder"
)

func main() {
	result := builder.NewArtBuilder().
		AddText("HI").
		SetStyle("bold").
		Build()

	fmt.Print(result)
}
