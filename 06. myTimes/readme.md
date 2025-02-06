# Time Operations in Go

## Table of Contents
- [Basic Time Operations](#basic-time-operations)
- [Time Formatting](#time-formatting)
- [Time Zones](#time-zones)
- [Time Calculations](#time-calculations)
- [Parsing and Conversion](#parsing-and-conversion)
- [Duration Operations](#duration-operations)
- [Calendar Functions](#calendar-functions)
- [Best Practices](#best-practices)
- [Performance Considerations](#performance-considerations)

## Basic Time Operations

### Current Time
```go
// Get current time
now := time.Now()

// Get components
year := now.Year()      // Get year
month := now.Month()    // Get month
day := now.Day()        // Get day
hour := now.Hour()      // Get hour
minute := now.Minute()  // Get minute
second := now.Second()  // Get second
nano := now.Nanosecond() // Get nanosecond
```

### Unix Timestamps
```go
// Get Unix timestamps
unix := now.Unix()        // Seconds since epoch
unixMilli := now.UnixMilli()  // Milliseconds since epoch
unixNano := now.UnixNano()    // Nanoseconds since epoch

// Convert Unix timestamp to Time
timeFromUnix := time.Unix(unix, 0)
```

## Time Formatting

### Standard Formats
```go
now := time.Now()

// Reference time: Mon Jan 2 15:04:05 MST 2006 (2006-01-02 15:04:05)
formats := map[string]string{
    "YYYY-MM-DD":     "2006-01-02",
    "DD-MM-YYYY":     "02-01-2006",
    "HH:mm:ss":       "15:04:05",
    "hh:mm:ss AM/PM": "03:04:05 PM",
    "Full DateTime":   "Monday, January 2, 2006 15:04:05",
}

// Usage
formatted := now.Format("2006-01-02 15:04:05")
```

### Predefined Formats
```go
// Built-in formats
now.Format(time.RFC3339)     // "2024-02-06T15:04:05Z07:00"
now.Format(time.RFC822)      // "06 Feb 24 15:04 MST"
now.Format(time.Kitchen)     // "3:04PM"
now.Format(time.Stamp)       // "Feb 6 15:04:05"
```

## Time Zones

### Time Zone Operations
```go
// Get local and UTC time
localTime := time.Now()
utcTime := localTime.UTC()

// Load specific timezone
nyc, _ := time.LoadLocation("America/New_York")
nycTime := localTime.In(nyc)

// Create time in specific zone
specificTime := time.Date(
    2024, time.February, 6,
    15, 04, 05, 0,
    nyc,
)
```

### Common Time Zone Conversions
```go
// Convert between zones
zones := []string{
    "America/Los_Angeles",
    "Europe/London",
    "Asia/Tokyo",
    "Australia/Sydney",
}

for _, zone := range zones {
    loc, _ := time.LoadLocation(zone)
    zoneTime := time.Now().In(loc)
    fmt.Printf("%s: %s\n", zone, zoneTime.Format("15:04:05 MST"))
}
```

## Time Calculations

### Addition and Subtraction
```go
now := time.Now()

// Add durations
tomorrow := now.Add(24 * time.Hour)
nextWeek := now.Add(7 * 24 * time.Hour)
twoHoursLater := now.Add(2 * time.Hour)

// Subtract durations
yesterday := now.Add(-24 * time.Hour)
lastWeek := now.Add(-7 * 24 * time.Hour)

// Add specific periods
nextMonth := now.AddDate(0, 1, 0)  // Add 1 month
nextYear := now.AddDate(1, 0, 0)   // Add 1 year
```

### Time Comparison
```go
time1 := time.Now()
time2 := time1.Add(time.Hour)

// Compare times
isBefore := time1.Before(time2)
isAfter := time1.After(time2)
isEqual := time1.Equal(time2)

// Calculate duration between times
duration := time2.Sub(time1)
```

## Parsing and Conversion

### String to Time
```go
// Parse time string
timeStr := "2024-02-06 15:04:05"
parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)

// Parse with location
loc, _ := time.LoadLocation("America/New_York")
parsedWithZone, err := time.ParseInLocation(
    "2006-01-02 15:04:05",
    timeStr,
    loc,
)
```

### Time to String
```go
// Convert time to string with custom format
timeString := time.Now().Format("Monday, January 2, 2006 at 3:04 PM")

// ISO 8601 format
isoString := time.Now().Format("2006-01-02T15:04:05Z07:00")
```

## Duration Operations

### Creating Durations
```go
// Standard durations
hour := time.Hour
minute := time.Minute
second := time.Second
millisecond := time.Millisecond

// Custom durations
customDuration := 2*time.Hour + 30*time.Minute
```

### Duration Calculations
```go
duration := 2 * time.Hour + 30 * time.Minute

// Get components
hours := duration.Hours()       // Get hours as float64
minutes := duration.Minutes()   // Get minutes as float64
seconds := duration.Seconds()   // Get seconds as float64
```

## Calendar Functions

### Month Operations
```go
now := time.Now()

// Get first day of month
firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

// Get last day of month
lastDay := firstDay.AddDate(0, 1, -1)

// Get days in month
daysInMonth := lastDay.Day()
```

### Week Operations
```go
// Get ISO week
year, week := now.ISOWeek()

// Get weekday
weekday := now.Weekday()  // Returns time.Weekday (0 = Sunday)

// Get day of year
dayOfYear := now.YearDay()
```

## Best Practices

### 1. Time Comparison
```go
// Correct way to compare times
if time1.Equal(time2) {
    // Times are equal
}

// Compare with consideration for timezone
if time1.UTC().Equal(time2.UTC()) {
    // Times are equal regardless of zone
}
```

### 2. Error Handling
```go
// Always check errors when parsing
if t, err := time.Parse(format, timeStr); err != nil {
    // Handle error
} else {
    // Use parsed time
}

// Check location errors
if loc, err := time.LoadLocation(zone); err != nil {
    // Handle error
}
```

### 3. Zone Management
```go
// Use UTC for storage and internal operations
storageTime := time.Now().UTC()

// Convert to local only for display
displayTime := storageTime.Local()
```

## Performance Considerations

### 1. Time Creation
```go
// Reuse time.Location objects
loc, _ := time.LoadLocation("America/New_York")
// Use loc multiple times instead of loading for each operation
```

### 2. Parsing Optimization
```go
// Cache parsed layout for repeated use
layout := "2006-01-02 15:04:05"
for _, timeStr := range timeStrings {
    t, _ := time.Parse(layout, timeStr)
    // Process t
}
```

### 3. Batch Processing
```go
// Process time operations in batches
times := make([]time.Time, 0, batchSize)
for i := 0; i < total; i += batchSize {
    // Process batch
}
```

This documentation provides a comprehensive overview of time operations in Go, focusing on practical implementations and best practices. The examples and guidelines can be adapted based on specific requirements while maintaining performance and reliability.