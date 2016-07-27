package main

import "fmt"

func main() {
	fmt.Print("Enter the number of feet:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * .3048
	fmt.Println(output, "meters")
}
