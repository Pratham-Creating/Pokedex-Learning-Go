package pokeapi

import (
	"encoding/json"
	"net/http"
)

var nameUrl = "https://pokeapi.co/api/v2/location-area/"

func ResponseNames(name string) (respNames, error) {

	NewUrl := nameUrl + name

	req, err := http.NewRequest("GET", NewUrl, nil)
	if err != nil {
		return respNames{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return respNames{}, err
	}
	defer res.Body.Close()

	var resName respNames
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resName); err != nil {
		return respNames{}, err
	}

	return resName, nil
}
