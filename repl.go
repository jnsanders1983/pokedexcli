package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var url = "https://pokeapi.co/api/"
var apiVersion = "v2/"
var location_area = "location-area/"

type LocationAreaResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands             map[string]cliCommand
	locationAreaResponse LocationAreaResponse
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

		err := callback.callback(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandMap(c *config) error {
	resp, err := http.Get(url + apiVersion + location_area)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &c.locationAreaResponse)
	if err != nil {
		return err
	}

	for _, locationArea := range c.locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(c *config) error {
	if c.locationAreaResponse.Previous == "" {
		return errors.New("you're on the first page")
	}
	return nil
}
