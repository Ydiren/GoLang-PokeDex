package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
)

const (
	DefaultApiLocationsUri = "https://pokeapi.co/api/v2/location/"
)

type PokeData struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (p *PokeData) CopyFrom(other *PokeData) {
	p.Count = other.Count
	p.Next = other.Next
	p.Previous = other.Previous
	p.Results = slices.Clone(other.Results)
}

func GetNextLocations(data *PokeData) (*PokeData, error) {
	nextUri := data.Next

	data, err := getDataFromApi(nextUri)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getDataFromApi(url string) (*PokeData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to close response body with error '%s'", err))
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Response failed with code '%d' and body '%s'", resp.StatusCode, resp.Body))
	}

	newData := &PokeData{}
	err = json.Unmarshal(body, &newData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return newData, nil
}

func GetPreviousLocations(data *PokeData) (*PokeData, error) {
	var previousUri string
	if data.Previous == nil {
		return nil, errors.New("cannot go to previous locations. You're already at the first location in the list")
	} else {
		previousUri = *data.Previous
	}

	data, err := getDataFromApi(previousUri)

	if err != nil {
		return nil, err
	}

	return data, nil
}
