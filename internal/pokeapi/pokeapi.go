package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ydiren/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const (
	DefaultApiLocationsUri = "https://pokeapi.co/api/v2/location/"
)

var cache = pokecache.NewCache(5 * time.Minute)

func (data *PokeData) GetNextLocations() error {
	nextUri := data.Next

	err := data.getDataFromApi(nextUri)
	if err != nil {
		return err
	}

	return nil
}

func (data *PokeData) GetPreviousLocations() error {
	var previousUri string
	if data.Previous == nil {
		return errors.New("Cannot go to previous locations. You're already at the first location in the list")
	} else {
		previousUri = *data.Previous
	}

	err := data.getDataFromApi(previousUri)

	if err != nil {
		return err
	}

	return nil
}

func (data *PokeData) getDataFromApi(url string) error {
	val, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to close response body with error '%s'", err))
		return err
	}

	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Response failed with code '%d' and body '%s'", resp.StatusCode, resp.Body))
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cache.Add(url, body)

	return nil
}
