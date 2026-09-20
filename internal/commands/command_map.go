package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jnsanders1983/pokedexcli/internal/pokeapi"
)

type LocationAreaResponse struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(c *Config, args []string) error {
	var mapURL string
	var body []byte

	if c.LocationAreaResponse.Next != nil {
		mapURL = *c.LocationAreaResponse.Next
	} else {
		mapURL = pokeapi.PokeAPIBaseURL + pokeapi.PokeAPIPath + pokeapi.PokeAPILocationAreas
	}

	cachedBody, ok := c.Cache.Get(mapURL)
	if !ok {
		resp, err := http.Get(mapURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		c.Cache.Add(mapURL, body)
	} else {
		fmt.Println("Cache hit for URL:", mapURL)
		body = cachedBody
	}

	err := json.Unmarshal(body, &c.LocationAreaResponse)
	if err != nil {
		return err
	}

	for _, locationArea := range c.LocationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(c *Config, args []string) error {
	var mapURL string
	var body []byte

	if c.LocationAreaResponse.Previous == nil {
		return errors.New("you're on the first page")
	}
	mapURL = *c.LocationAreaResponse.Previous

	cachedBody, ok := c.Cache.Get(mapURL)
	if !ok {
		resp, err := http.Get(mapURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		c.Cache.Add(mapURL, body)
	} else {
		fmt.Println("Cache hit for URL:", mapURL)
		body = cachedBody
	}

	err := json.Unmarshal(body, &c.LocationAreaResponse)
	if err != nil {
		return err
	}

	for _, locationArea := range c.LocationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
