package main

import (
	"bufio"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nPlease enter 1 to lookup your weather, 2 to retrieve previous searches with your username or 3 to quit the program")
		scanner.Scan()

		if scanner.Text() == "3" {
			fmt.Print("\nGoodbye")
			break
		}

		if scanner.Text() == "2" {
			fmt.Println("Please enter your username to retrieve your data")
			usernameToSearch := scanner.Scan()
			var product []Product
			db.Where("username <> ?", usernameToSearch).Find(&product)
			retrievedValue1, returnedValue2 := db.Table("products").Select("COALESCE(username,?)", usernameToSearch).Rows()
			fmt.Println(retrievedValue1, returnedValue2)

		}

		if scanner.Text() == "1" {

			fmt.Println("\nPlease enter your username (make one up if you dont have one)")
			scanner.Scan()
			var userNameInput string = scanner.Text()

			fmt.Println("\nPlease enter a password")
			scanner.Scan()
			var passwordInput string = scanner.Text()

			fmt.Println("\nPlease enter your address")
			scanner.Scan()
			var rawAddressInput string = scanner.Text()

			returnedLongitude, returnedLatitude, returnedLocality := getUserLocation(rawAddressInput)

			currentTemp, currentWeatherDescription, AltLocationName := getWeather(returnedLatitude, returnedLongitude)

			fmt.Sprint(AltLocationName)

			fmt.Printf("\nThe Current tempreature in %v is %v and the weather is currently %v \n\n", returnedLocality, currentTemp, currentWeatherDescription)

			fmt.Println("\n\n", userNameInput, passwordInput, rawAddressInput, currentWeatherDescription, currentTemp)

			db.Create(&Product{Username: userNameInput, Password: passwordInput, Address: rawAddressInput, WeatherDescription: currentWeatherDescription, Tempreature: currentTemp})

			fmt.Println("Weather Saved in DB")
		}
	}
}
