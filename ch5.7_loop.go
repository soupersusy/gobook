package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// if i is Divisable by 3 AND 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "blob")
		}
	}
}
