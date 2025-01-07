package main

import "fmt"

func main() {
	collatz_conjecture(25)
}

func collatz_conjecture(num int) {
	for i := 1; num > 1; i += 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
		fmt.Printf("%d. %d\n", i, num)
	}
}
