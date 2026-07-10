package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func getWorkDayStatus(day Weekday) string {
	switch day {
	case Sunday, Monday, Tuesday, Wednesday:
		return "Office is open"
	case Thursday:
		return "Half day open"
	case Friday, Saturday:
		return "Off day"
	default:
		return "Invalid Day"

	}
}

type OfficeStatus string

const (
	StatusOpen    OfficeStatus = "open"
	StatusClosed  OfficeStatus = "closed"
	StatusHalfDay OfficeStatus = "half_day"
)

func main() {
	fmt.Println(getWorkDayStatus(Friday))
}
