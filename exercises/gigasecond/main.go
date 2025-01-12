package main

import (
	"fmt"
	"time"
)

var input_date string = "1997-03-23"
var layout string = "2006-01-02"

func date_parser(date_str string) time.Time {
	parsed_date, err := time.Parse(layout, date_str)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return parsed_date
}

func date_formatter(new_date time.Time) string {
	formatted_date := new_date.Format(layout)
	return formatted_date
}

func add_gigasecond(input_date string) {
	seconds := int64(1000000000)
	parsed_date := date_parser(input_date)
	new_date := parsed_date.Add(time.Duration(seconds) * time.Second)
	formatted_date := date_formatter(new_date)
	fmt.Println("New date after adding 1,000,000,000 seconds:", formatted_date)
}

func main() {
	add_gigasecond(input_date)
}
