package p

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var cardinals = []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}

// ForecastAPIV1Struct : Main Forecast Struct V1
type ForecastAPIV1Struct struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		LastUpdated time.Time `json:"last_updated"`
		TempC       float64   `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		WindKph  float64 `json:"wind_kph"`
		WindDir  string  `json:"wind_dir"`
		PrecipMm float64 `json:"precip_mm"`
	} `json:"current"`
	Forecast struct {
		Forecastday []ForecastdayAPIV1Struct `json:"forecastday"`
	} `json:"forecast"`
}

// ForecastdayAPIV1Struct : Main Forecastday Struct V1
type ForecastdayAPIV1Struct struct {
	Date time.Time `json:"date"`
	Day  struct {
		MaxtempC  float64 `json:"maxtemp_c"`
		MintempC  float64 `json:"mintemp_c"`
		AvgtempC  float64 `json:"avgtemp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"day"`
}

// LocationAPIV1Struct is the main Forecastday Struct V1
type LocationAPIV1Struct struct {
	Name   string `json:"name"`
	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"latLng"`
}

// Forecast contains the forecast functions
type Forecast interface {
	GetForecastFromName(string) (ForecastAPIV1Struct, error)
}

// Location contains the location functions
type Location interface {
	GetLocationFromName(string) (LocationAPIV1Struct, error)
}

// fetchAPI fetch an api and convert the api structure to th Forecast API V1 structure
func fetchAPI(url string, timeout int) ([]byte, error) {
	var err error

	spaceClient := http.Client{
		Timeout: time.Second * time.Duration(timeout), // Maximum on secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "forecast-test")

	res, err := spaceClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DegToCardinal convert degress to short string cardinal
func DegToCardinal(deg float64) (string, error) {
	val := int((deg / 22.5) + .5)
	if val >= len(cardinals) || val < 0 {
		return "", fmt.Errorf("Non degree value")
	}
	return cardinals[val], nil
}
