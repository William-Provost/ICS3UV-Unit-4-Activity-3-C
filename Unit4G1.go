// Author: Mr Coxall
// Version: 1.0.0
// Date: 2025-01-01
// Fileoverview: This program finds average temperature, using arrays.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	// Array of 5 floating-point numbers
	var temperature [5]float64
	var scanner = bufio.NewScanner(os.Stdin)
	var sum float64 = 0

	// get five temperatures from user
	for counter := 0; counter < 5; counter++ {
		fmt.Printf("Enter temperature %d: ", counter+1)
		scanner.Scan()
		userInput := scanner.Text()
		
		// Parse the user input string into a float64
		oneTemperature, _ := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
		temperature[counter] = oneTemperature
	}

	// calculate sum using a for loop
	for counter := 0; counter < len(temperature); counter++ {
		sum += temperature[counter]
	}

	// calculate average to two decimal places
	// Need to cast len(temperature) to float64 for division
	average := sum / float64(len(temperature))

	// display all temperatures using a for loop
	fmt.Println("\nThe temperatures are:")
	for counter := 0; counter < len(temperature); counter++ {
		// %d is integer, so we use %f for float
		fmt.Printf("Temperature %d: %.0f°C\n", counter+1, temperature[counter]) 
	}

	// display the average
	// %.2f formats the float to two decimal places
	fmt.Printf("\nThe average temperature is: %.2f°C\n", average)

	fmt.Println("\nDone.")
}
