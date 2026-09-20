package commands

import (
	"github.com/jnsanders1983/pokedexcli/internal/pokeapi"
	"github.com/jnsanders1983/pokedexcli/internal/pokecache"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

type Config struct {
	Commands             map[string]Command
	LocationAreaResponse LocationAreaResponse
	Cache                *pokecache.Cache
	Pokedex              map[string]pokeapi.Pokemon
}

func NewConfig(cache *pokecache.Cache) *Config {
	return &Config{
		Commands: map[string]Command{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Callback:    commandExit,
			},
			"help": {
				Name:        "help",
				Description: "Display help information",
				Callback:    commandHelp,
			},
			"map": {
				Name: "map",
				Description: "Displays the names of 20 location areas in the Pokemon world. " +
					"Each subsequent call to the map command will display the " +
					"next 20 location areas",
				Callback: commandMap,
			},
			"mapb": {
				Name: "mapb",
				Description: "Displays the names of the previous 20 location areas in the " +
					"Pokemon world. Each subsequent call to the mapb command will display the " +
					"previous 20 location areas",
				Callback: commandMapBack,
			},
			"explore": {
				Name:        "explore",
				Description: "Explore a specific location area by name",
				Callback:    commandExplore,
			},
			"catch": {
				Name:        "catch",
				Description: "Attempt to catch a Pokemon by name",
				Callback:    commandCatch,
			},
			"inspect": {
				Name:        "inspect",
				Description: "Inspect a caught Pokemon",
				Callback:    commandInspect,
			},
			"pokedex": {
				Name:        "pokedex",
				Description: "List all caught Pokemon",
				Callback:    commandPokedex,
			},
		},
		LocationAreaResponse: LocationAreaResponse{},
		Cache:                cache,
		Pokedex:              make(map[string]pokeapi.Pokemon),
	}
}
