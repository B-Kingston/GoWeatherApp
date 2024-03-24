package main

import(
  "fmt"
  "net/http"
  "io"
  "encoding/json"
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



func getWeather(latitudeInput, longitudeInput float64) (float64, string, string) {

  //Calls MapBox to get locality of address, OpenWeatherMap is innacurate
  locality := getLocality(longitudeInput, latitudeInput)

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5//weather?lat=%v&lon=%v&units=metric&lang=en&appid=%s", latitudeInput, longitudeInput, openWeatherApiKey)
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

  return storedWeatherResponse.Main.Temp, storedWeatherResponse.Weather[0].Description, locality
}
