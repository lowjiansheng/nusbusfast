/*
 * This calls the BUS API to retrieve bus information.
 */
package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BASE_URL = "https://nextbus.comfortdelgro.com.sg/eventservice.svc/"

type busStopResult struct {
	busStops struct {
		stops []struct {
			caption   string  `json:"caption"`
			latitude  float64 `json:"latitude"`
			longitude float64 `json:"longitude"`
			name      string  `json:"name"`
		} `json:"busstops"`
	} `json:"BusStopsResult"`
}

func ListBusStops() []byte {
	route := "BusStops"
	fmt.Println(BASE_URL + route)
	resp, err := http.Get(BASE_URL + route)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// The API returned garbage data. Should return error.
	}
	var busRes busStopResult
	jsonErr := json.Unmarshal(body, &busRes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return body
}
