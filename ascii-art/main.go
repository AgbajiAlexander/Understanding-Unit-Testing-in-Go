package main

import (
	"fmt"
)

func StringToArt(s string) string {
	// minimal local implementation to avoid the external package dependency
	// replace with the real conversion logic as needed
	return s
}

func main() {
	examples := []string{"0", "1", "12", "0123456789", "1\n2"}

	for _, input := range examples {
		fmt.Printf("Input: %q\n", input)
		fmt.Println(StringToArt(input))
		fmt.Println("---")
	}
}
