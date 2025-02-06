package main

import (
	"fmt"
	"time"
)

func main() {
	// Initialize and clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[1;36m=== Time Operations in Go ===\033[0m")
	fmt.Println("\033[1;33mExploring time manipulation and formatting\033[0m")
	time.Sleep(1 * time.Second)

	// 1. Current Time Operations
	fmt.Println("\n\033[1;32m1. CURRENT TIME\033[0m")
	fmt.Println("---------------")
	now := time.Now()
	fmt.Printf("Current Time:     %v\n", now)
	fmt.Printf("UTC Time:         %v\n", now.UTC())
	fmt.Printf("Unix Timestamp:   %v\n", now.Unix())
	fmt.Printf("Unix Nano:        %v\n", now.UnixNano())
	fmt.Printf("Unix Milli:       %v\n", now.UnixMilli())

	// 2. Time Components
	fmt.Println("\n\033[1;32m2. TIME COMPONENTS\033[0m")
	fmt.Println("-----------------")
	fmt.Printf("Year:             %d\n", now.Year())
	fmt.Printf("Month:            %s (%d)\n", now.Month(), now.Month())
	fmt.Printf("Day:              %d\n", now.Day())
	fmt.Printf("Hour:             %d\n", now.Hour())
	fmt.Printf("Minute:           %d\n", now.Minute())
	fmt.Printf("Second:           %d\n", now.Second())
	fmt.Printf("Nanosecond:       %d\n", now.Nanosecond())
	fmt.Printf("Weekday:          %s\n", now.Weekday())
	fmt.Printf("Day of Year:      %d\n", now.YearDay())
	// Get ISO week number
	_, week := now.ISOWeek()
	fmt.Printf("Week Number:      %d\n", week)

	// 3. Time Formatting
	fmt.Println("\n\033[1;32m3. TIME FORMATTING\033[0m")
	fmt.Println("-----------------")
	fmt.Printf("Standard:         %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("With Timezone:    %v\n", now.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("Custom Format 1:  %v\n", now.Format("Monday, January 2, 2006"))
	fmt.Printf("Custom Format 2:  %v\n", now.Format("3:04 PM - Jan 2, 2006"))
	fmt.Printf("ISO 8601:         %v\n", now.Format("2006-01-02T15:04:05Z07:00"))
	fmt.Printf("RFC3339:          %v\n", now.Format(time.RFC3339))
	fmt.Printf("Kitchen Format:   %v\n", now.Format(time.Kitchen))

	// 4. Time Zone Operations
	fmt.Println("\n\033[1;32m4. TIME ZONES\033[0m")
	fmt.Println("-------------")

	// Load various time zones
	locations := []string{
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Australia/Sydney",
		"Pacific/Auckland",
	}

	for _, loc := range locations {
		if location, err := time.LoadLocation(loc); err == nil {
			localTime := now.In(location)
			fmt.Printf("%-16s %s\n", loc+":", localTime.Format("15:04:05 MST"))
		}
	}

	// 5. Time Calculations
	fmt.Println("\n\033[1;32m5. TIME CALCULATIONS\033[0m")
	fmt.Println("-------------------")

	// Future times
	fmt.Println("\nFuture Times:")
	fmt.Printf("Tomorrow:         %v\n", now.Add(24*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Printf("Next Week:        %v\n", now.Add(7*24*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Printf("Next Month:       %v\n", now.AddDate(0, 1, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Next Year:        %v\n", now.AddDate(1, 0, 0).Format("2006-01-02 15:04:05"))

	// Past times
	fmt.Println("\nPast Times:")
	fmt.Printf("Yesterday:        %v\n", now.Add(-24*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Printf("Last Week:        %v\n", now.Add(-7*24*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Printf("Last Month:       %v\n", now.AddDate(0, -1, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Last Year:        %v\n", now.AddDate(-1, 0, 0).Format("2006-01-02 15:04:05"))

	// 6. Duration Calculations
	fmt.Println("\n\033[1;32m6. DURATION CALCULATIONS\033[0m")
	fmt.Println("----------------------")

	// Create some dates for comparison
	futureDate := now.AddDate(0, 0, 10) // 10 days from now
	pastDate := now.AddDate(0, 0, -10)  // 10 days ago

	// Calculate durations
	futureDuration := futureDate.Sub(now)
	pastDuration := now.Sub(pastDate)

	fmt.Printf("Future Duration:   %v\n", futureDuration)
	fmt.Printf("Past Duration:     %v\n", pastDuration)
	fmt.Printf("In Hours:         %.2f\n", futureDuration.Hours())
	fmt.Printf("In Minutes:       %.2f\n", futureDuration.Minutes())
	fmt.Printf("In Seconds:       %.2f\n", futureDuration.Seconds())

	// 7. Time Parsing
	fmt.Println("\n\033[1;32m7. TIME PARSING\033[0m")
	fmt.Println("---------------")

	timeFormats := []string{
		"2006-01-02 15:04:05",
		"Jan 2, 2006 3:04 PM",
		"2006-01-02T15:04:05Z07:00",
		"Monday, January 2, 2006",
		"15:04:05",
	}

	sampleTime := "2024-02-06 15:04:05"
	fmt.Printf("Parsing: %s\n", sampleTime)

	for _, format := range timeFormats {
		if parsed, err := time.Parse(format, sampleTime); err == nil {
			fmt.Printf("With format %-25s: %v\n", format, parsed)
		}
	}

	// 8. Time Comparison
	fmt.Println("\n\033[1;32m8. TIME COMPARISON\033[0m")
	fmt.Println("------------------")

	time1 := now
	time2 := now.Add(time.Hour)
	time3 := now.Add(-time.Hour)

	fmt.Printf("Time1:            %v\n", time1.Format("15:04:05"))
	fmt.Printf("Time2:            %v\n", time2.Format("15:04:05"))
	fmt.Printf("Time3:            %v\n", time3.Format("15:04:05"))
	fmt.Printf("Time1 before Time2: %v\n", time1.Before(time2))
	fmt.Printf("Time1 after Time3:  %v\n", time1.After(time3))
	fmt.Printf("Time1 equal Time1:  %v\n", time1.Equal(time1))

	// 9. Calendar Calculations
	fmt.Println("\n\033[1;32m9. CALENDAR CALCULATIONS\033[0m")
	fmt.Println("-----------------------")

	// First and last day of current month
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Printf("First of Month:    %v\n", firstOfMonth.Format("2006-01-02"))
	fmt.Printf("Last of Month:     %v\n", lastOfMonth.Format("2006-01-02"))
	fmt.Printf("Days in Month:     %d\n", lastOfMonth.Day())

	// Days until end of year
	lastOfYear := time.Date(currentYear, 12, 31, 0, 0, 0, 0, now.Location())
	daysLeft := lastOfYear.Sub(now).Hours() / 24
	fmt.Printf("Days left in year: %.0f\n", daysLeft)

	// Program completion
	fmt.Println("\n\033[1;36m=== Program Complete ===\033[0m")
	fmt.Println("\033[1;33mPress Enter to exit...\033[0m")
	fmt.Scanln()
}
