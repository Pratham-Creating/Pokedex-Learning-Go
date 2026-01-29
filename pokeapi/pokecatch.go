package pokeapi

import (
	"encoding/json"
	"net/http"
)

var Url = "https://pokeapi.co/api/v2/pokemon/"

func ReponseCatch(name string) (RespCatch, error) {

	newUrl := Url + name

	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return RespCatch{}, err

	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return RespCatch{}, err
	}

	defer res.Body.Close()

	var respcatch RespCatch
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&respcatch); err != nil {
		return RespCatch{}, err
	}

	return respcatch, nil
}
