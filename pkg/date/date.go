// Package date provides utility functions for working with time and date operations in Go.
// It offers a simplified interface for adding, subtracting, formatting, parsing, and
// creating time values.
//
// The package also defines common format constants and duration helpers for time calculations,
// including seconds, minutes, hours, days, weeks, months, and years.
//
// Example usage:
//
//	now := date.Now()
//	fmt.Println("Now:", now.Format(date.Ddmmyyyyhhmmss))
//
//	tomorrow := date.Tomorrow()
//	fmt.Println("Tomorrow:", tomorrow)
//
//	future := date.IsYears.AddDate(5)
//	fmt.Println("Five years later:", future)
package date

import (
	"log"
	"time"
)

const (
	// Ddmmyyyy defines the date format as "02/01/2006" (day/month/year).
	Ddmmyyyy = "02/01/2006"
	// Mmddyyyy defines the date format as "01/02/2006" (month/day/year).
	Mmddyyyy = "01/02/2006"
	// Yyyymmdd defines the date format as "2006/01/02" (year/month/day).
	Yyyymmdd = "2006/01/02"
	// Ddmmyyyyhhmmss defines the full date and time format as "02/01/2006 03:04:05".
	Ddmmyyyyhhmmss = "02/01/2006 03:04:05"
	// Hhmmss defines the time format as "03:04:05" (hour:minute:second).
	Hhmmss = "03:04:05"

	// Second represents the duration of one second.
	Second = 1 * time.Second
	// Minute represents the duration of 60 seconds.
	Minute = 60 * time.Second
	// Hour represents the duration of 12 hours.
	Hour = 12 * time.Hour
	// Day represents the duration of 24 hours.
	Day = 24 * time.Hour
	// Month represents an approximate duration of 30 days.
	Month = 30 * Day
	// Year represents an approximate duration of 365 days.
	Year = 365 * Day
	// Week represents the duration of 7 days.
	Week = 7 * Day
)

// Date defines the type of date increment operation (days, months, or years).
type Date uint8

const (
	// IsDays indicates that the increment will be performed in days.
	IsDays Date = iota
	// IsMonths indicates that the increment will be performed in months.
	IsMonths
	// IsYears indicates that the increment will be performed in years.
	IsYears
)

// AddDate adjusts the current time by adding the specified value based on the Date type:
// days, months, or years. The base reference is always the current time.
func (d Date) AddDate(value int) time.Time {
	switch d {
	case IsDays:
		return time.Now().AddDate(0, 0, value)
	case IsMonths:
		return time.Now().AddDate(0, value, 0)
	case IsYears:
		return time.Now().AddDate(value, 0, 0)
	default:
		return time.Now()
	}
}

// Equals checks weather two time.Time values a and b are equal using the Equal method.
func Equals(a, b time.Time) bool {
	return a.Equal(b)
}

// Add adds a given time.Duration value to the provided time.Time,
// returning a new time.Time instance.
func Add(value time.Duration, current *time.Time) time.Time {
	return current.Add(value)
}

// Sub subtracts a given time.Duration value from the provided time.Time,
// returning a new time.Time instance.
func Sub(value time.Duration, current *time.Time) time.Time {
	return Add(-value, current)
}

// Format formats the given time.Time using the specified layout string.
// Example: date.Format(date.Ddmmyyyy, &time.Now())
func Format(format string, current *time.Time) string {
	return current.Format(format)
}

// Born creates a new time.Time based on the provided year, month, and day.
// The time component is set to midnight (00:00:00) in UTC.
func Born(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// Now returns the current local time.
func Now() time.Time {
	return time.Now()
}

// Today returns the current date truncated to midnight (00:00:00).
func Today() time.Time {
	return time.Now().Truncate(Day)
}

// Tomorrow returns the date corresponding to tomorrow (00:00:00).
func Tomorrow() time.Time {
	return Today().Add(Day)
}

// Yesterday returns the date corresponding to yesterday (00:00:00).
func Yesterday() time.Time {
	return Today().Add(-Day)
}

// Parse converts a string into a time.Time using the specified layout format.
// If parsing fails, the program exits with log.Fatal.
func Parse(format string, value string) time.Time {
	res, err := time.Parse(format, value)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// Seconds returns a time.Duration representing the specified number of seconds.
func Seconds(seconds int) time.Duration {
	return Second * time.Duration(seconds)
}

// Minutes returns a time.Duration representing the specified number of minutes.
func Minutes(minutes int) time.Duration {
	return Minute * time.Duration(minutes)
}

// Hours returns a time.Duration representing the specified number of hours.
func Hours(hours int) time.Duration {
	return Hour * time.Duration(hours)
}

// Days returns a time.Duration representing the specified number of days.
func Days(days int) time.Duration {
	return Day * time.Duration(days)
}

// Months returns an approximate time.Duration representing the specified number of months.
func Months(months int) time.Duration {
	return Month * time.Duration(months)
}

// Years returns an approximate time.Duration representing the specified number of years.
func Years(years int) time.Duration {
	return Year * time.Duration(years)
}
