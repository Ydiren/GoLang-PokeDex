package pokeapi

import (
	"encoding/json"
	"errors"
)

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

	pokemon := make([]Pokemon, len(location.PokemonEncounters))
	for i := 0; i < len(location.PokemonEncounters); i++ {
		pokemon[i] = location.PokemonEncounters[i].Pokemon
	}

	return pokemon, nil
}
