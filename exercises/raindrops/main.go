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

func rain_sounds(nums ...int) {
	for _, num := range nums {
		fmt.Println(rain_sound(num))
	}
}

func main() {
	numbers := []int{12, 42, 1, 43, 23, 13, 432, 54, 28, 30, 34, 64, 24, 653, 97, 2231, 54, 32, 54, 1245, 0}

	fmt.Println(rain_sound(34))
	rain_sounds(numbers...)
}
