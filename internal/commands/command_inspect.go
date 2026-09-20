package commands

import (
	"fmt"
)

func commandInspect(c *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: inspect <pokemon_name>")
	}

	pokemonName := args[0]
	pokemon, caught := c.Pokedex[pokemonName]
	if !caught {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	return nil
}
