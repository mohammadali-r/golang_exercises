// task: https://exercism.org/tracks/go/exercises/two-fer

package main

import "fmt"

func offer(name string){
	if name == ""{
		fmt.Println("One for you, one for me")
	}else{
		fmt.Printf("One for %s, one for me\n", name)
	}
}

func main() {
	name := ""
	offer(name)
}
