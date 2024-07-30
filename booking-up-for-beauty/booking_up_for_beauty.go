package booking

/*
The `time.Parse` function parses strings into values of type `Time`.
Go has a special way of how you define the layout you expect for the parsing.
You need to write an example of the layout using the values from this special timestamp:
				`Mon Jan 2 15:04:05 -0700 MST 2006`.
*/

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	return getTime(layout, date)
}

func getTime(layout, date string) time.Time {
	t, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t := getTime(layout, date)

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon (>= 12:00 and < 18:00).
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t := getTime(layout, date)

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t := getTime(layout, date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday().String(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	const anniversaryMonth = 9
	const anniversaryDay = 15

	return time.Date(time.Now().Year(), anniversaryMonth, anniversaryDay, 0, 0, 0, 0, time.UTC)
}
