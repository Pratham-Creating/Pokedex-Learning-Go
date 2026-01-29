package pokeapi

//Making this struct using the information from the API
//https://pokeapi.co/api/v2/location-area/

//-------------------------------------------------------
//Make sure to use capital letter if you need to export it
//-------------------------------------------------------

type RespLocations struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"` //Using * as this can be empty too1
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
