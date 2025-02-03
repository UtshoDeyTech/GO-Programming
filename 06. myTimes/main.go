package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	// Current time (your original code)
	presentTime := time.Now()
	fmt.Println("\n=== Current Time ===")
	fmt.Println("Current time:", presentTime)
	fmt.Println("Formatted:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Created date (your original code)
	createDate := time.Date(2020, time.June, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println("\n=== Created Date ===")
	fmt.Println("Created date:", createDate)
	fmt.Println("Formatted:", createDate.Format("01-02-2006 Monday"))

	// Different Time Formats
	fmt.Println("\n=== Time Formatting ===")
	fmt.Println("YYYY-MM-DD:", presentTime.Format("2006-01-02"))
	fmt.Println("DD-MM-YYYY:", presentTime.Format("02-01-2006"))
	fmt.Println("12 Hour:", presentTime.Format("03:04:05 PM"))
	fmt.Println("24 Hour:", presentTime.Format("15:04:05"))
	fmt.Println("Custom Format:", presentTime.Format("Monday, January 2, 2006 at 3:04 PM"))

	// Time Components
	fmt.Println("\n=== Time Components ===")
	fmt.Printf("Year: %d\n", presentTime.Year())
	fmt.Printf("Month: %s\n", presentTime.Month())
	fmt.Printf("Day: %d\n", presentTime.Day())
	fmt.Printf("Hour: %d\n", presentTime.Hour())
	fmt.Printf("Minute: %d\n", presentTime.Minute())
	fmt.Printf("Second: %d\n", presentTime.Second())
	fmt.Printf("Weekday: %s\n", presentTime.Weekday())
	fmt.Printf("Day of Year: %d\n", presentTime.YearDay())

	// Time Zone Operations
	fmt.Println("\n=== Time Zones ===")
	localTime := presentTime
	utcTime := presentTime.UTC()
	fmt.Printf("Local Time: %v\n", localTime.Format("15:04:05 MST"))
	fmt.Printf("UTC Time: %v\n", utcTime.Format("15:04:05 MST"))

	// Load specific timezone
	estLocation, _ := time.LoadLocation("America/New_York")
	estTime := presentTime.In(estLocation)
	fmt.Printf("New York Time: %v\n", estTime.Format("15:04:05 MST"))

	// Time Calculations
	fmt.Println("\n=== Time Calculations ===")
	futureTime := presentTime.Add(time.Hour * 24) // Add 24 hours
	fmt.Printf("Tomorrow: %v\n", futureTime.Format("2006-01-02 15:04:05"))

	pastTime := presentTime.Add(-time.Hour * 24) // Subtract 24 hours
	fmt.Printf("Yesterday: %v\n", pastTime.Format("2006-01-02 15:04:05"))

	// Duration Calculations
	fmt.Println("\n=== Duration Calculations ===")
	duration := futureTime.Sub(presentTime)
	fmt.Printf("Duration between times: %v\n", duration)

	// Time Parsing
	fmt.Println("\n=== Time Parsing ===")
	timeStr := "2024-02-03 15:04:05"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Printf("Parsed time: %v\n", parsedTime)
	}

	// Working with Unix Timestamps
	fmt.Println("\n=== Unix Timestamps ===")
	unixTime := presentTime.Unix()
	unixNano := presentTime.UnixNano()
	fmt.Printf("Unix Timestamp: %v\n", unixTime)
	fmt.Printf("Unix Nano Timestamp: %v\n", unixNano)

	// Convert Unix timestamp back to Time
	timeFromUnix := time.Unix(unixTime, 0)
	fmt.Printf("Time from Unix: %v\n", timeFromUnix.Format("2006-01-02 15:04:05"))

	// Time Comparison
	fmt.Println("\n=== Time Comparison ===")
	time1 := time.Now()
	time2 := time1.Add(time.Hour)
	fmt.Printf("time1 before time2: %v\n", time1.Before(time2))
	fmt.Printf("time1 after time2: %v\n", time1.After(time2))
	fmt.Printf("times equal: %v\n", time1.Equal(time2))

	// Sleep and Timer Example
	fmt.Println("\n=== Sleep and Timer ===")
	fmt.Println("Starting short sleep...")
	time.Sleep(2 * time.Second)
	fmt.Println("Sleep finished!")

	// Timer example
	timer := time.NewTimer(1 * time.Second)
	<-timer.C
	fmt.Println("Timer finished!")

	// Date Difference
	fmt.Println("\n=== Date Difference ===")
	date1 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC)
	daysDiff := date2.Sub(date1).Hours() / 24
	fmt.Printf("Days between dates: %.0f\n", daysDiff)
}
