package main

import (
	"fmt"
	"os"
	"time"
)

var location = "Heilbronn"

const (
	APIkey          = "c3733bad37c94ddba6a32813250504&q="
	days            = "6"
	redisAddr       = "localhost:6379"
	redisPassword   = ""
	redisDB         = 0
	cacheExpiration = 5 * time.Minute
)

func main() {

	if len(os.Args) >= 2 {
		switch {
		case os.Args[1] == "help":
			openHelp()
			return
		case os.Args[1] == "ext":
			openExtendedMenu()
			return
		default:
			if len(os.Args) >= 1 {
				location = os.Args[1]
			}
		}
	}
	weather, err := getWeather(location)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		return
	}

	printLocalForecast(weather)
}

func printLocalForecast(weather Weather) {
	location := weather.Location
	current := weather.Current
	hours := weather.Forecast.Forecastday[0].Hour

	printLoc := fmt.Sprintf("%s, %s: %.1f°C -  %s -  Wind: %.1f km/h -->  %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
		current.WindKph,
		current.WindDir)

	fmt.Printf("\033[32m") //green
	fmt.Printf("\033[1m")  //bold
	fmt.Print(printLoc)
	fmt.Printf("\033[0m") // reset color

	for _, hour := range hours {
		date := time.Unix(int64(hour.TimeEpoch), 0)
		if date.Before(time.Now()) {
			continue
		}
		if date.After(time.Now().Add(12 * time.Hour)) {
			break
		}
		message := fmt.Sprintf("%s: %.1f°C\t %s\t Chances of rain: %d\t Wind: %.1f km/h --> %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.Condition.Text,
			hour.ChanceOfRain,
			hour.WindKph,
			hour.WindDir,
		)

		switch {
		case hour.Condition.Text == "Sunny":
			fmt.Printf("\033[33m") //yellow
			fmt.Print(message)
			fmt.Printf("\033[0m")
		case hour.ChanceOfRain > 10:
			fmt.Printf("\033[34m") //blue
			fmt.Print(message)
			fmt.Printf("\033[0m")
		default:
			fmt.Print(message)
		}
	}
}
