package main

import "fmt"

func main() {
	//for statement allows us to repeat a list of statements (a block) multiple times
	// var 1 eq 1; if 1 is less eq 10 then do the if then add 1 then loop to for
	for i := 1; i <= 10; i++ {
		//if - way of choosing to do different things based on a condition
		// i %2==0 - var divided by 2 with no remainder
		if i%2 == 0 {
			fmt.Println(i, "even")
			// optional else - if the if condition is false, move to else
		} else {
			fmt.Println(i, "odd")
		}
	}
}
