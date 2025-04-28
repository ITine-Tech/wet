package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func openExtendedMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the extended menu. Press 'q' to quit. \nYou can access more weather options.\nPlease choose your location first:")
	fmt.Scan(&location)
	if location == "q" || location == "Q" {
		fmt.Println("Exiting.")
		os.Exit(0)
	}
	fmt.Println("Options:")
	fmt.Println("<a>:\t Display astronomical information")
	fmt.Println("<1> to <5>: Display weather forecast for the upcoming days")

	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "q", "Q":
			fmt.Println("Exiting")
			os.Exit(0)
		case "1", "2", "3", "4", "5":
			printTomorrowForecast(input)
			return
		case "a", "A":
			printAstroInfo()
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func printAstroInfo() {

	weather, err := getWeather(location)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}
	moon := weather.Forecast.Forecastday[0].Astro.IsMoonUp

	fmt.Printf("Astronomical information for %s:\n", location)
	fmt.Printf("* Sunrise: %s\n", weather.Forecast.Forecastday[0].Astro.Sunrise)
	fmt.Printf("* Sunset: %s\n", weather.Forecast.Forecastday[0].Astro.Sunset)
	fmt.Printf("* Moonrise: %s\n", weather.Forecast.Forecastday[0].Astro.Moonrise)
	fmt.Printf("* Moonset: %s\n", weather.Forecast.Forecastday[0].Astro.Moonset)
	fmt.Printf("* Moon Phase: %s\n", weather.Forecast.Forecastday[0].Astro.MoonPhase)
	fmt.Printf("* Moon Illumination: %d%%\n", weather.Forecast.Forecastday[0].Astro.MoonIllumination)
	switch moon {
	case 1:
		fmt.Println("* Moon is up")
	default:
		fmt.Println("* Moon is not up")
	}
}

func printTomorrowForecast(input string) {
	day, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Didn't get correct input.")
		return
	}
	tomorrow := time.Now().AddDate(0, 0, day)
	found := false

	weather, err := getWeather(location)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}
	for _, day := range weather.Forecast.Forecastday {
		{
			if day.Date == tomorrow.Format("2006-01-02") {
				found = true
				fmt.Printf("\033[1m")
				fmt.Printf("\033[32m")
				fmt.Printf("\n%s -  %s\n", location, tomorrow.Format("Monday, 2 January 2006"))
				fmt.Printf("\033[0m")
				fmt.Printf("* Max temperature: %.1f°C\n", day.Day.MaxtempC)
				fmt.Printf("* Min temperature: %.1f°C\n", day.Day.MintempC)
				fmt.Printf("* Condition: %s\n", day.Day.Condition.Text)
				fmt.Printf("* Wind: %.1f km/h\n", day.Day.MaxwindKph)
				fmt.Printf("* Chances of Rain: %d%%\n", day.Day.DailyChanceOfRain)
				break
			}
		}

	}
	if !found {
		fmt.Println("Weather forecast not found.")
	}
}
