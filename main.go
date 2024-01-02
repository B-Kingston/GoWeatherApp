package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type weatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type geoCoder []struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
}

func getUserLocation(rawAddress string) string {
	fmt.Println(rawAddress, "formattedLatLon")
	geoCoderApiKey := os.Getenv("GeoCoderAPI")
	formattedLatLon := strings.ReplaceAll(rawAddress, " ", "+")

	fmt.Println(formattedLatLon, "formattedLatLon")

	url := fmt.Sprintf("https://geocode.maps.co/search?q=%s&api_key=%s", formattedLatLon, geoCoderApiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	geocodeResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var geocodeResponse = geoCoder{}
	errGeocodeResponse := json.Unmarshal(geocodeResponseBody, &geocodeResponse)
	log.Print(errGeocodeResponse)

	latLon := geocodeResponse[0].Lat

	return latLon
}

func main() {
	var rawAddressInput string
	openWeatherMapApiKey := os.Getenv("GeoCoderAPI")
	fmt.Print(openWeatherMapApiKey)
	fmt.Println("What is your address:")
	fmt.Scanln(&rawAddressInput)
	//fmt.Println(rawAddressInput)
	returnedLatLon := getUserLocation(rawAddressInput)
	fmt.Println(returnedLatLon)

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5//weather?id=2060771&units=metric&lang=en&appid=%s", openWeatherMapApiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	WeatherResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var storedWeatherResponse = weatherResponse{}
	errWeatherResponse := json.Unmarshal(WeatherResponseBody, &storedWeatherResponse)
	log.Print(errWeatherResponse)

	fmt.Println("The weather in swanview is currently", storedWeatherResponse.Weather[0].Description, "and a temp of", storedWeatherResponse.Main.Temp)

}
