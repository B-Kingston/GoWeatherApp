package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type AutoGenerated struct {
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

func getUserLocation() {
	var userAddress string
	var formattedAddress string

	fmt.Println("What is your address:")
	fmt.Scan(&userAddress)
	formattedAddress = strings.ReplaceAll(userAddress, " ", "+")

	url := "https://geocode.maps.co/search?q=&api_key=6593bf0164db0782915133phm4da85d"
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	getUserLocation()

	url := "http://api.openweathermap.org/data/2.5//weather?id=2060771&units=metric&lang=en&appid=0c0f275ebb573dc5d3d95ba639702b0d"
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		//fmt.Println(err)
		return
	}

	//test := body
	var test2 = AutoGenerated{}
	err2 := json.Unmarshal(body, &test2)
	log.Print(err2)
	// log.Printf("%+v", test2.Weather)
	// fmt.Printf(string(test))
	// stringToCheck := test2.Weather
	// convertedString := strconv.Itoa(stringToCheck)

	fmt.Println("The weather in swanview is currently", test2.Weather[0].Description, "and a temp of", test2.Main.Temp)
	//if strings.Contains(stringToCheck,"rain"){

	//}
}