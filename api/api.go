package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/cjsampson/weather-service/app/conf"
)

type WeatherAPI interface {
	Get(_ context.Context, req *http.Request) (WeatherResponse, error)
}

func NewWeatherAPI(config conf.WeatherConfig) WeatherAPI {
	return &weatherAPI{
		BaseURL: config.BaseURL,
		APIKey:  config.APIKey,
	}
}

type weatherAPI struct {
	BaseURL string
	APIKey  string
}

func (w *weatherAPI) Get(_ context.Context, req *http.Request) (WeatherResponse, error) {
	var wresp WeatherResponse

	params, err := validateRequest(req)
	if err != nil {
		log.Printf("validateRequest failed - %v\n", err.Error())
	}

	resp, err := http.Get(w.constructURI(w.APIKey, params))
	if err != nil {
		log.Printf("http.get failed - %v\n", err.Error())

		return wresp, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("io.ReadAll failed - %v\n", err)
		return wresp, err
	}

	if err = json.Unmarshal(body, &wresp); err != nil {
		log.Printf("json.Unmarshal failed - %v\n", err)
		return wresp, err
	}

	return wresp, nil
}

// WeatherParams - container for the validated latitude && longitude
type WeatherParams struct {
	Latitude  string
	Longitude string
}

// validateRequest - validation function to check for query param existence
func validateRequest(req *http.Request) (WeatherParams, error) {
	var (
		lat    = req.URL.Query().Get("lat")
		long   = req.URL.Query().Get("lon")
		params = WeatherParams{}
	)

	if lat == "" || long == "" {
		return params, fmt.Errorf("latitude or longitude invalid - lat: %v long: %v\n", lat, long)
	}

	params.Latitude = lat
	params.Longitude = long

	return params, nil
}

// constructURI - manually add to preserve order
func (w *weatherAPI) constructURI(apiKey string, ps WeatherParams) string {
	encodedApiKey := url.QueryEscape(apiKey)
	encodedLat := url.QueryEscape(ps.Latitude)
	encodedLong := url.QueryEscape(ps.Longitude)

	queryString := fmt.Sprintf("lat=%s&lon=%s&appid=%s", encodedLat, encodedLong, encodedApiKey)

	return fmt.Sprintf("%s?%s", w.BaseURL, queryString)
}

type WeatherResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Rain       Rain      `json:"rain,omitempty"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level,omitempty"`
	GrndLevel int     `json:"grnd_level,omitempty"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust,omitempty"`
}

type Rain struct {
	OneHour float64 `json:"1h,omitempty"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
