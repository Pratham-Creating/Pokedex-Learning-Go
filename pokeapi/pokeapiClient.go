package pokeapi

import (
	"encoding/json"
	"net/http"
)

var baseUrl = "https://pokeapi.co/api/v2"

func ListLocations(pageUrl *string) (RespLocations, error) { //Here also we use a pointer as the function will also work if we have a nil value in the parameter.
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil) //It just helps to get the request so we can use that variable to recieve a response
	if err != nil {
		return RespLocations{}, err
	}

	client := &http.Client{} //This is acutally responsible to send the request over internet
	res, err := client.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	defer res.Body.Close()

	var resp RespLocations

	decoder := json.NewDecoder(res.Body) //Decoding the stream of bytes to GO struct
	if err := decoder.Decode(&resp); err != nil {
		return RespLocations{}, err
	}
	return resp, nil

}
