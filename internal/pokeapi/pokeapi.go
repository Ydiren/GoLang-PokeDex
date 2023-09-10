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
	PokemonDetailsUri      = "https://pokeapi.co/api/v2/pokemon/"
)

var cache = pokecache.NewCache(5 * time.Minute)

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

	if resp.StatusCode == 404 {
		return nil, errors.New(fmt.Sprintf("Not found!"))
	}
	if resp.StatusCode > 399 {
		return nil, errors.New(fmt.Sprintf("Response failed with code '%d' and body '%s'", resp.StatusCode, resp.Body))
	}

	cache.Add(url, body)

	return body, nil
}
