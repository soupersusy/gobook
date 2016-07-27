package main

import "fmt"

func main() {
	i := 1
	// for statement allows us to repeat a list of statements (a block) multiple times
	//if i is less than or eq to 10
	for i <= 10 {
		//println
		fmt.Println(i)
		//then add 1 and loopback to the for statement
		i = i + 1
	}
}
