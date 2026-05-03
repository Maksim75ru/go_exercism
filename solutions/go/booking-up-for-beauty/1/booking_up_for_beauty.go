package booking

import (
    "time"
    "fmt"
    "log"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)

    if err != nil {
        log.Printf("Something goes wrong: %v", err) 
        panic("Sometign goes wrong")
    }
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)

    if err != nil {
        log.Printf("Something goes wrong: %v", err) 
        panic("Sometign goes wrong ")
    }

    return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)

    if err != nil {
        log.Printf("Something goes wrong: %v", err) 
        panic("Sometign goes wrong ")
    }
    
    hour := t.Hour()
    minute := t.Minute()
    
    return 12 <= hour && minute >= 0 && hour < 18 && minute <= 59
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)

    if err != nil {
        log.Printf("Something goes wrong: %v", err) 
        panic("Sometign goes wrong ")
    }
    formattedTime := t.Format("Monday, January 2, 2006, at 15:04") // string

    return fmt.Sprintf("You have an appointment on %v.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    now := time.Now()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
