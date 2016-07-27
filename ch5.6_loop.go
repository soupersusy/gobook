package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "Fizz3")
		} else if i%5 == 0 {
			fmt.Println(i, "ugh5")
		} else {
			fmt.Println(i)
		}
	}
}
