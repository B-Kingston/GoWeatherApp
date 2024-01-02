# Weather Application

***Written by ChatGPT lul***

This Go program retrieves weather information for a given location using the OpenWeatherMap API.

## Prerequisites

Before running the program, make sure you have the following prerequisites:

- Go installed on your machine.
- An API key from OpenWeatherMap. You can obtain one by signing up here.

## Installation

1. Clone the repository:
    
    bashCopy code
    
    `git clone [repository_url] cd [repository_directory]`
    
2. Replace `[API_KEY]` in the `url` variable inside the `main` function with your OpenWeatherMap API key.
    
3. Build and run the program:
    
    bashCopy code
    
    `go build main.go ./main`
    

## Usage

1. Enter your address when prompted.
2. The program will use the address to obtain latitude and longitude information.
3. It will then fetch and display the current weather information for the specified location.

## Configuration

You can customize the program by modifying the following parameters:

- **OpenWeatherMap API Key:** Replace `[API_KEY]` in the `url` variable inside the `main` function with your API key.

## Acknowledgments

- Weather data is provided by [OpenWeatherMap](https://openweathermap.org/).