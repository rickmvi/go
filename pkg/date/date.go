package date

import (
	"time"
)

const (
	// --- Format Constants ---
	Ddmmyyyy       = "02/01/2006"
	Mmddyyyy       = "01/02/2006"
	Yyyymmdd       = "2006/01/02"
	Ddmmyyyyhhmmss = "02/01/2006 03:04:05"
	Hhmmss         = "03:04:05"

	// --- Duration Constants ---
	Second = time.Second
	Minute = time.Minute
	Hour   = time.Hour // Corrigido para 1 hora
	Day    = 24 * time.Hour
	Week   = 7 * Day
)

// Now returns the current local time as a value of type time.Time.
func Now() time.Time {
	return time.Now()
}

// Today returns the current date truncated to the start of the day (00:00:00).
func Today() time.Time {
	// time.Truncate(Day) é mais idiomático do que time.Now().Truncate(Day)
	return time.Now().Truncate(Day)
}

// Tomorrow returns the date for the next day by adding 24 hours to the current day's start (00:00:00).
func Tomorrow() time.Time {
	return Today().Add(Day)
}

// Yesterday returns the date corresponding to the start of the previous day (00:00:00).
func Yesterday() time.Time {
	return Today().Add(-Day)
}

// Born creates a new time.Time object representing the specified year, month, and day at midnight in UTC.
func Born(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// Parse parses a formatted string and returns the corresponding time value or an error if the format is invalid.
func Parse(format string, value string) (time.Time, error) {
	return time.Parse(format, value)
}

// --- Funções de Manipulação de Tempo e Duração (Ponteiros removidos) ---

// Equals compares two time.Time values and returns true if they are equal, otherwise returns false.
func Equals(a, b time.Time) bool {
	return a.Equal(b)
}

// Add adds a specified time.Duration to the provided time.Time and returns the resulting time.
func Add(value time.Duration, current time.Time) time.Time {
	return current.Add(value)
}

// Sub subtracts a specified time.Duration from the provided time.Time and returns the resulting time.
func Sub(value time.Duration, current time.Time) time.Time {
	return Add(-value, current)
}

// Format formata a time.Time usando o layout string especificado.
func Format(format string, current time.Time) string {
	return current.Format(format)
}

// AddDays adds the specified number of days to the given time and returns the resulting time value.
func AddDays(current time.Time, days int) time.Time {
	return current.AddDate(0, 0, days)
}

// AddMonths adds the specified number of months to the given time and returns the resulting time.
func AddMonths(current time.Time, months int) time.Time {
	return current.AddDate(0, months, 0)
}

// AddYears adds the specified number of years to the given time and returns the resulting time value.
func AddYears(current time.Time, years int) time.Time {
	return current.AddDate(years, 0, 0)
}

// Seconds converts an integer value representing seconds into a time.Duration type in seconds.
func Seconds(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

// Minutes converts an integer value representing minutes into a time.Duration type in minutes.
func Minutes(minutes int) time.Duration {
	return time.Duration(minutes) * time.Minute
}

// Hours converts the given number of hours into a time.Duration object representing the equivalent duration in hours.
func Hours(hours int) time.Duration {
	return time.Duration(hours) * time.Hour
}

// Days converts the provided number of days into a time.Duration representing the equivalent duration in hours.
func Days(days int) time.Duration {
	return time.Duration(days) * Day
}
