package pokeapi

import "encoding/json"

func GetPokemonDetails(pokemonName *string) (*PokemonDetails, error) {
	pokemonUri := PokemonDetailsUri + *pokemonName
	body, err := getDataFromApi(pokemonUri)
	if err != nil {
		return nil, err
	}

	pokemonDetails := PokemonDetails{}
	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		return nil, err
	}

	return &pokemonDetails, nil
}
