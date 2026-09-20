package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/jnsanders1983/pokedexcli/internal/pokeapi"
)

func commandCatch(c *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon_name>")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonURL := pokeapi.PokeAPIBaseURL + pokeapi.PokeAPIPath + pokeapi.PokeAPIPokemon + pokemonName + "/"

	body, ok := c.Cache.Get(pokemonURL)
	if !ok {
		response, err := http.Get(pokemonURL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusNotFound {
			return fmt.Errorf("pokemon '%s' not found", pokemonName)
		}
		if response.StatusCode > 299 {
			return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
		}

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		c.Cache.Add(pokemonURL, body)
	}

	var pokemon pokeapi.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	chance := catchChance(pokemon.BaseExperience)
	if rand.Float64() < chance {
		c.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon.Name)

	return nil
}

func catchChance(baseExperience int) float64 {
	chance := 1 - float64(baseExperience)/1000
	if chance < 0.1 {
		return 0.1
	}
	if chance > 0.95 {
		return 0.95
	}
	return chance
}
