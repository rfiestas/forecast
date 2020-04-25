// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"log"
)

// providers init
func providers() Forecast {
	var locationProvider Location
	var forecastProvider Forecast
	var err error

	// Init Location Provider
	locationProvider, err = NewMapquest()
	if err != nil {
		log.Fatal(err)
	}
	// Init Forecast Provider
	forecastProvider, err = NewOpenweather(locationProvider)
	if err != nil {
		log.Fatal(err)
	}
	return forecastProvider
}

// getQueryKey take a http request url query key, assign default value when not exist.
func getQueryKey(r *http.Request, key string, failure string) string {
	var value string
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		value = failure
	} else {
		value = string(keys[0])
	}
	return url.QueryEscape(value)
}

// ForecastAPIV1 Forecast call
func ForecastAPIV1(w http.ResponseWriter, r *http.Request) {
	var location string
	// Init Providers
	provider := providers()

	// Fetch Forecast
	location = getQueryKey(r, "location", "Barcelona")
	forecast, err := provider.GetForecastFromName(url.QueryEscape(location))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Convert to Json and print
	res, err := json.Marshal(forecast)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, string(res))
}

// GetIndex returns the index.html
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var location string
	location = getQueryKey(r, "location", "Barcelona")
	data, err := ioutil.ReadFile("templates/indexV1.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Header().Set("Cache-Control", "max-age=86400, public")
	fmt.Fprint(w, strings.Replace(string(data), "#LOCATION#", location, -1))
	return
}

// GetRobots returns the robots.txt
func GetRobots(w http.ResponseWriter, r *http.Request) {
	var data string
	data = "User-agent: *\nDisallow: \n"
	fmt.Fprint(w, data)
	return
}
