package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/encapsulationandembedding/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(1994)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
