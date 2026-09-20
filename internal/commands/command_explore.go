package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jnsanders1983/pokedexcli/internal/pokeapi"
)

type ExploreResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *Config, args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location area name (e.g., explore pastoria-city-area)")
	}
	areaName := args[0]

	mapURL := pokeapi.PokeAPIBaseURL + pokeapi.PokeAPIPath + pokeapi.PokeAPILocationAreas + areaName

	var body []byte

	cachedBody, ok := c.Cache.Get(mapURL)
	if !ok {
		fmt.Printf("Exploring %s...\n", areaName)
		resp, err := http.Get(mapURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			return fmt.Errorf("location area '%s' not found", areaName)
		}
		if resp.StatusCode > 299 {
			return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
		}

		// Use '=' to avoid shadowing any nested logic contexts
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		c.Cache.Add(mapURL, body)
	} else {
		fmt.Printf("Exploring %s (from cache)...\n", areaName)
		body = cachedBody
	}

	var exploreData ExploreResponse
	err := json.Unmarshal(body, &exploreData)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	if len(exploreData.PokemonEncounters) == 0 {
		fmt.Println(" - No wild Pokemon found in this area.")
		return nil
	}

	for _, encounter := range exploreData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
