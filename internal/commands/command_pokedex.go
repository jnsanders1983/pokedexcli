package commands

import (
	"fmt"
)

func commandPokedex(c *Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
