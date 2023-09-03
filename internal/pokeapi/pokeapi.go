package pokeapi

import (
	"errors"
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const (
	DefaultApiLocationsUri = "https://pokeapi.co/api/v2/location-area/"
)

var cache = pokecache.NewCache(5 * time.Minute)

func (data *PokeLocations) GetNextLocations() error {
	nextUri := data.Next

	return data.getLocationsData(nextUri)
}

func (data *PokeLocations) GetPreviousLocations() error {
	var previousUri string
	if data.Previous == nil {
		return errors.New("Cannot go to previous locations. You're already at the first location in the list")
	} else {
		previousUri = *data.Previous
	}

	return data.getLocationsData(previousUri)
}

func (data *PokeLocations) getLocationsData(previousUri string) error {
	body, err := getDataFromApi(previousUri)
	if err != nil {
		return err
	}

	err = data.parseData(body)
	if err != nil {
		return err
	}

	cache.Add(previousUri, body)
	return nil
}

func getDataFromApi(url string) ([]byte, error) {
	val, ok := cache.Get(url)
	if ok {
		return val, nil
	}

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

	return body, nil
}
