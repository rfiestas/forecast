package p 

import (
    "time"
)

// ForecastAPIV1Struct : Main Forecast Struct V1 
type ForecastAPIV1Struct struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
   	    LastUpdated time.Time `json:"last_updated"`
		TempC       float64   `json:"temp_c"`
		Condition  struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		WindKph     float64   `json:"wind_kph"`
		WindDir     string    `json:"wind_dir"`
		PrecipMm    float64   `json:"precip_mm"`
    } `json:"current"`
	Forecast struct {
		Forecastday []ForecastdayAPIV1Struct `json:"forecastday"`
	} `json:"forecast"`
}

// ForecastdayAPIV1Struct : Main Forecastday Struct V1
type ForecastdayAPIV1Struct struct {
    Date time.Time `json:"date"`
    Day  struct {
        MaxtempC float64 `json:"maxtemp_c"`
        MintempC float64 `json:"mintemp_c"`
        AvgtempC float64 `json:"avgtemp_c"`
        Condition  struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
        } `json:"condition"`
    } `json:"day"`
}