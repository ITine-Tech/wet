package main

import "fmt"

func openHelp() {
	fmt.Println("Help menu")
	fmt.Printf("* <city>: Provide a city when executing in order to receive local weather information, i.e. 'wet Berlin'\n")
	fmt.Printf("* <ext>: Use 'ext' to open the extended menu, which includes additional weather information\n")
}
