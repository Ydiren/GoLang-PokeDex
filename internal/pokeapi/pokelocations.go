package pokeapi

import (
	"encoding/json"
	"errors"
)

func (locations *PokeLocations) parseData(data []byte) error {
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	return nil
}

func (locations *PokeLocations) GetNextLocations() error {
	nextUri := locations.Next

	return locations.getLocationsData(nextUri)
}

func (locations *PokeLocations) GetPreviousLocations() error {
	var previousUri string
	if locations.Previous == nil {
		return errors.New("Cannot go to previous locations. You're already at the first location in the list")
	} else {
		previousUri = *locations.Previous
	}

	return locations.getLocationsData(previousUri)
}

func (locations *PokeLocations) getLocationsData(previousUri string) error {
	body, err := getDataFromApi(previousUri)
	if err != nil {
		return err
	}

	err = locations.parseData(body)
	if err != nil {
		return err
	}

	cache.Add(previousUri, body)
	return nil
}
