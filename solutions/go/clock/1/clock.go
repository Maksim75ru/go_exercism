package clock

import (
    "fmt"
)

// Define the Clock type here.
type Clock struct {
    h int
    m int
}


func New(h, m int) Clock {
   	totalHours := h + m/60
    minutes := m % 60
    
    if minutes < 0 {
        minutes += 60
        totalHours--
    }

    hours := totalHours % 24
    if hours < 0 {
        hours += 24
    }
    
	return Clock{
        h: hours,
        m: minutes,
    }
}

func (c Clock) Add(m int) Clock {
	subHour := c.h + (c.m + m) / 60
    subMin := (c.m + m) % 60
    return New(subHour, subMin)
}

func (c Clock) Subtract(m int) Clock {
	totalHours := c.h
    minutes := c.m - m
    
	if minutes < 0 {
        minutes += 60
        totalHours--
    }
    
    return New(totalHours, minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%2.2d:%2.2d", c.h, c.m)
}
