package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz3")
			// optional else - if the if condition is false, move to else
		} else {
			fmt.Println("ugh")
		}
	}
}
