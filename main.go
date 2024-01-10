package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

type geoCoder struct {
	Type     string   `json:"type"`
	Query    []string `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			Accuracy string `json:"accuracy"`
			MapboxID string `json:"mapbox_id"`
		} `json:"properties"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Address string `json:"address"`
		Context []struct {
			ID        string `json:"id"`
			MapboxID  string `json:"mapbox_id"`
			Text      string `json:"text"`
			Wikidata  string `json:"wikidata,omitempty"`
			ShortCode string `json:"short_code,omitempty"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}

func getUserLocation(rawAddress string) (float64, float64, string) {

	geoCoderApiKey := os.Getenv("MapBoxAPI")
	formattedLatLon := strings.ReplaceAll(rawAddress, " ", "%20")

	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s&limit=1", formattedLatLon, geoCoderApiKey)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, 0, "err"
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, 0, "err"
	}
	defer res.Body.Close()

	geocodeResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, 0, "err"
	}

	var geocodeResponse = geoCoder{}
	json.Unmarshal(geocodeResponseBody, &geocodeResponse)

	longitude := geocodeResponse.Features[0].Geometry.Coordinates[0]
	latitude := geocodeResponse.Features[0].Geometry.Coordinates[1]
	locality := geocodeResponse.Features[0].Context[1].Text
	return longitude, latitude, locality
}

func getWeather(latitudeInput, longitudeInput float64) (float64, string, string) {
	openWeatherMapApiKey := os.Getenv("OpenWeatherAPI")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5//weather?lat=%v&lon=%v&units=metric&lang=en&appid=%s", latitudeInput, longitudeInput, openWeatherMapApiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, "", ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, "", ""
	}
	defer res.Body.Close()

	WeatherResponseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, "", ""
	}

	var storedWeatherResponse = weatherResponse{}
	json.Unmarshal(WeatherResponseBody, &storedWeatherResponse)
	//fmt.Println(storedWeatherResponse.Main.Temp, storedWeatherResponse.Weather[0].Description, "!!!!!!!!!!!!!!!!!")
	return storedWeatherResponse.Main.Temp, storedWeatherResponse.Weather[0].Description, storedWeatherResponse.Name
}

func newLatLon(originalLon, originalLat float64) (string, string, string, string) {

	altLatNorth := fmt.Sprintf("%.9g", originalLat-0.13)
	altLonEast := fmt.Sprintf("%.9g", originalLon+0.15)
	altLatSouth := fmt.Sprintf("%.9g", originalLat+0.13)
	altLonWest := fmt.Sprintf("%.9g", originalLon-0.15)

	altLatNorth1, err1 := strconv.ParseFloat(altLatNorth, 64)
	altLonEast1, err2 := strconv.ParseFloat(altLonEast, 64)
	altLatSouth1, err3 := strconv.ParseFloat(altLatSouth, 64)
	altLonWest1, err4 := strconv.ParseFloat(altLonWest, 64)

	fmt.Sprint(err1, err2, err3, err4) //clears error for unused err's

	altWeatherNorthTemp, altWeatherNorthDescription, altNorthLocationName := getWeather(altLatNorth1, originalLon)
	combinedNorth := fmt.Sprintf("The current weather in %s is %s with a tempreature of %v", altNorthLocationName, altWeatherNorthDescription, altWeatherNorthTemp)

	time.Sleep(2 * time.Second)

	altWeatherEastTemp, altWeatherEastDescription, altEastLocationName := getWeather(originalLat, altLonEast1)
	combinedEast := fmt.Sprintf("The current weather in %s is %s with a tempreature of %v", altEastLocationName, altWeatherEastDescription, altWeatherEastTemp)

	time.Sleep(2 * time.Second)

	altWeatherSouthTemp, altWeatherSouthDescription, altSouthLocationName := getWeather(altLatSouth1, originalLon)
	combinedSouth := fmt.Sprintf("The current weather in %s is %s with a tempreature of %v", altSouthLocationName, altWeatherSouthDescription, altWeatherSouthTemp)

	time.Sleep(2 * time.Second)

	altWeatherWestTemp, altWeatherWestDescription, altWestLocationName := getWeather(originalLat, altLonWest1)
	combinedWest := fmt.Sprintf("The current weather in %s is %s with a tempreature of %v", altWestLocationName, altWeatherWestDescription, altWeatherWestTemp)

	return combinedNorth, combinedEast, combinedSouth, combinedWest
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nPlease enter your desired location, or type Q to exit")
		scanner.Scan()
		if scanner.Text() == "Q" || scanner.Text() == "q" {
			fmt.Print("\nGoodbye")
			break
		} else {
			rawAddressInput := scanner.Text()

			returnedLongitude, returnedLatitude, returnedLocality := getUserLocation(rawAddressInput)

			currentTemp, currentWeatherDescription, AltLocationName := getWeather(returnedLatitude, returnedLongitude)

			fmt.Sprint(AltLocationName)

			altWeatherNorth, altWeatherEast, altWeatherSouth, altWeatherWest := newLatLon(returnedLongitude, returnedLatitude)
			fmt.Println("The weather north of your location is", altWeatherNorth)
			fmt.Println("The weather east of your location is", altWeatherEast)
			fmt.Println("The weather south of your location is", altWeatherSouth)
			fmt.Println("The weather west of your location is", altWeatherWest)

			fmt.Printf("\nThe Current tempreature in %v is %v and the weather is currently %v \n\n", returnedLocality, currentTemp, currentWeatherDescription)

		}
	}
}
