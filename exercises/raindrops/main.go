// task: https://exercism.org/tracks/go/exercises/raindrops

package main

import (
	"fmt"
	"strconv"
)

func rain_sound(num int) string {
	res := ""

	if num%3 == 0 {
		res += "Pling"
	}
	if num%5 == 0 {
		res += "Plang"
	}
	if num%7 == 0 {
		res += "Plong"
	}
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		res += strconv.Itoa(num)
	}

	return res
}

//TODO create function for do this on range

func main() {
	fmt.Println(rain_sound(34))
}
