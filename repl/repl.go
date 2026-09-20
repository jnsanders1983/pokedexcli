package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jnsanders1983/pokedexcli/internal/commands"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func Start(c *commands.Config) {
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

		callback, exists := c.Commands[command]
		if !exists {
			fmt.Println("Unknown command:", command)
			continue
		}

		err := callback.Callback(c, cleanedInput[1:])
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
