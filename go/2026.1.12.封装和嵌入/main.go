package main

import (
	"lipcoder/calendar"
	"log"
)

func main() {
	data := calendar.Date{}
	err := data.SetYear(23)
	if err != nil {
		log.Fatal(err)
	}
	data.Settime()
}
