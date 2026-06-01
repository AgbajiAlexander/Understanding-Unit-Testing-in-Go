package main

import (
	"ascii_art_font/font"

	"fmt"
)

func main() {
	myFont := font.GenerateFont()

	for r := rune(32); r <= rune(126); r++ {
		fmt.Printf("Character: %q (ASCII %d)\n", r, r)
		lines := myFont[r]
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println() // Empty line between characters
	}
}
