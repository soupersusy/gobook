package main

import "fmt"

func main() {
	// fmt.Print
	fmt.Print("Enter a number:")
	// take input as decimal - real #
	var input float64
	// scanf reads user input - fills input with the # entered
	fmt.Scanf("%f", &input)
	//define output variable
	// ":=" is the same as "var output = input*2"
	output := input * 2
	fmt.Println(input, "times 2 = ", output)
}
