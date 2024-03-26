package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weather", getWeather)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func getWeather(r http.ResponseWriter, req *http.Request) {
	var (
		lat  = req.URL.Query().Get("longitude")
		long = req.URL.Query().Get("latitude")
		url  = "https://api.openweathermap.org/data/2.5/weather"
		// ?q=London,uk&APPID=049dec87e84ff02bcf0902e1ad44f1c8"
	)

	fmt.Printf("response: %v\n", resp)

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("http.Get failed - %v\n", err.Error())
	}
}
