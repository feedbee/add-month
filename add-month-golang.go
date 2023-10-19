package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date and time
	currentDate := time.Now()
	// Add one month to the current date
	futureDate := currentDate.AddDate(0, -1, 0)

	// Print the current date and the date one month from now
	fmt.Println("Current Date:", currentDate)
	fmt.Println("Date One Month from Now:", futureDate)

	fmt.Println("")

	// Set a specific date (e.g., October 31, 2023)
	specificDate := time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)
    // Add one month to the specific date
	futureDate2 := specificDate.AddDate(0, -1, 0)

	// Print the specific date
	fmt.Println("Specific Date:", specificDate)
	fmt.Println("Date One Month from the Specific Date:", futureDate2)
}
