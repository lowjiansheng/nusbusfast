package main

import (
	"fetcher"
	"net/http"
)

const PORT_NUM = ":8080"

func listBusStops(w http.ResponseWriter, r *http.Request) {
	body := fetcher.ListBusStops()
	w.Write(body)
}

func main() {
	http.HandleFunc("/listbusstops", listBusStops)
	if err := http.ListenAndServe(PORT_NUM, nil); err != nil {
		panic(err)
	}
}
