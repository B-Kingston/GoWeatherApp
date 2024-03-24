package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)
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

type forwardGeoCoder struct {
	Type     string    `json:"type"`
	Query    []float64 `json:"query"`
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

	// geoCoderApiKey := os.Getenv("MapBoxAPI")
	formattedAddress := strings.ReplaceAll(rawAddress, " ", "%20")

	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s&limit=1", formattedAddress, mapBoxApiKey)
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
