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
	DefaultApiLocationsUri = "https://pokeapi.co/api/v2/location-area/"
)

var cache = pokecache.NewCache(5 * time.Minute)

func GetPokemonAtLocation(locationName *string) ([]Pokemon, error) {
	if locationName == nil {
		return nil, errors.New("locationName cannot be nil")
	}

	locationUri := DefaultApiLocationsUri + *locationName
	body, err := getDataFromApi(locationUri)
	if err != nil {
		return nil, err
	}

	location := PokeSingleLocation{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}

	cache.Add(locationUri, body)

	pokemon := make([]Pokemon, len(location.PokemonEncounters))
	for i := 0; i < len(location.PokemonEncounters); i++ {
		pokemon[i] = location.PokemonEncounters[i].Pokemon
	}

	return pokemon, nil
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
