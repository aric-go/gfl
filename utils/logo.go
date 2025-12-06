package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// DisplayLogo displays the claude-ai ASCII art logo with purple color and extra spacing
func DisplayLogo() {
	// Set purple color (ANSI code for purple)
	fmt.Printf("\x1b[35m\x1b[1m")

	// Print the figure with ivrit font
	figure := figure.NewFigure("GFL", "smslant", true)
	figure.Print()

	// Reset color to default
	fmt.Printf("\x1b[0m")

	// Add extra spacing
	fmt.Println()
}
