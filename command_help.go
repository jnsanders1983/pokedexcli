package main

import (
	"fmt"
	"strings"
)

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	const nameWidth = 8
	const descWidth = 50

	for _, cmd := range c.commands {
		lines := wrapText(cmd.description, descWidth)
		fmt.Printf("%-*s%s\n", nameWidth, cmd.name, lines[0])
		for _, line := range lines[1:] {
			fmt.Printf("%-*s%s\n", nameWidth, "", line)
		}
	}

	return nil
}

func wrapText(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	var current string

	for _, word := range words {
		if len(current)+len(word)+1 > width {
			lines = append(lines, current)
			current = word
		} else if current == "" {
			current = word
		} else {
			current += " " + word
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
