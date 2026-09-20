package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func commandMap(c *config, args []string) error {
	var mapURL string
	var body []byte

	if c.locationAreaResponse.Next != nil {
		mapURL = *c.locationAreaResponse.Next
	} else {
		mapURL = PokeAPIBaseURL + PokeAPIPath + PokeAPILocationAreas
	}

	cachedBody, ok := c.cache.Get(mapURL)
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
		c.cache.Add(mapURL, body)
	} else {
		fmt.Println("Cache hit for URL:", mapURL)
		body = cachedBody
	}

	err := json.Unmarshal(body, &c.locationAreaResponse)
	if err != nil {
		return err
	}

	for _, locationArea := range c.locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapBack(c *config, args []string) error {
	var mapURL string
	var body []byte

	if c.locationAreaResponse.Previous == nil {
		return errors.New("you're on the first page")
	}
	mapURL = *c.locationAreaResponse.Previous

	cachedBody, ok := c.cache.Get(mapURL)
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
		c.cache.Add(mapURL, body)
	} else {
		fmt.Println("Cache hit for URL:", mapURL)
		body = cachedBody
	}

	err := json.Unmarshal(body, &c.locationAreaResponse)
	if err != nil {
		return err
	}

	for _, locationArea := range c.locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
