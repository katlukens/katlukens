package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {

	// Define your date range
	startDate, _ := time.Parse("1/2/2006", "5/22/2024")
	endDate, _ := time.Parse("1/2/2006", "6/4/2024")

	// Open the CSV file
	file, err := os.Open("AccountHistory.1.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read and discard the header row
	if _, err := reader.Read(); err != nil {
		fmt.Println("Error reading header:", err)
		return
	}

	// Read the CSV content
	for {
		record, err := reader.Read()
		if err != nil {
			break // Stop at EOF or other error
		}

		// Parse the date from the second column
		recordDate, err := time.Parse("1/2/2006", record[1])
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue // Skip rows with invalid dates
		}
		// Check if the date is within the range
		if recordDate.After(startDate) && recordDate.Before(endDate) {
			fmt.Println(record) // Print the row
		}
	}
}
