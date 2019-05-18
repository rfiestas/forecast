// Package p contains an HTTP Cloud Function.
package p

import (
	"net/http"
	"fmt"
	"os";
	"io/ioutil"
	"net/url"
)


func ForecastAPIV1(w http.ResponseWriter, r *http.Request) {
	var location string
	keys, ok := r.URL.Query()["location"]
	if !ok || len(keys[0]) < 1 {
		location = "Barcelona"
	}else{
		location = string(keys[0])
	}

	fmt.Fprint(w,GetForecastAPIV1(url.QueryEscape(location)))
	return
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("templates/index.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprint(w,string(data))
	return
}
func GetRobots(w http.ResponseWriter, r *http.Request) {
	var data string
	data = "User-agent: *\nDisallow: /\n"
	fmt.Fprint(w,data)
	return
}