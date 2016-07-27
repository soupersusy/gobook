package main

import "fmt"

func main() {
	fmt.Print("Enter the temp in Fahrenhiet:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5 / 9
	fmt.Println(input, "Fahrenhiet equals", output, "Celcius")
}
