# Weather App README

***Written by ChatGPT but looks accurate***

## Introduction

This GoLang application provides real-time weather information based on user-provided location data. It leverages the Mapbox Geocoding API to convert a user's address input into geographical coordinates (latitude and longitude). The obtained coordinates are then used to fetch weather data from the OpenWeatherMap API.

## Prerequisites

Before running the application, ensure you have the following:

- Go installed on your machine.
- Access to Mapbox API and OpenWeatherMap API keys. Set these keys as environment variables:
  - Mapbox API Key: `MapBoxAPI`
  - OpenWeatherMap API Key: `OpenWeatherAPI`

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your_username/your_repo.git
    cd your_repo
    ```

2. Set up environment variables:

    ```bash
    export MapBoxAPI=your_mapbox_api_key
    export OpenWeatherAPI=your_openweather_api_key
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

## Usage

1. The application will prompt you to enter your location.

2. Provide your address as input when prompted.

3. The application will display the current temperature and weather description for the specified location.

## Code Structure

- `main.go`: Contains the main logic for user interaction, geocoding, and weather data retrieval.

## Dependencies

- [Mapbox Geocoding API](https://docs.mapbox.com/api/search/geocoding/)
- [OpenWeatherMap API](https://openweathermap.org/api)

## Acknowledgments

- Mapbox and OpenWeatherMap for providing the APIs used in this application.