package main

import "fmt"

func main() {
	// var i =10, if less and 100, do if statement then add one a loop back to for
	for i := 10; i <= 100; i++ {
		//i divided by 3 remainder 0 -
		if i%3 == 0 {
			fmt.Println(i, "Divisable by 3")
		}
	}
}
