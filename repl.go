package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jnsanders1983/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error // <-- Must have '[]string' here!
}

type config struct {
	commands             map[string]cliCommand
	locationAreaResponse LocationAreaResponse
	cache                *pokecache.Cache
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		command := cleanedInput[0]

		callback, exists := c.commands[command]
		if !exists {
			fmt.Println("Unknown command:", command)
			continue
		}

		err := callback.callback(c, cleanedInput[1:])
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
