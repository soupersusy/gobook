//Write a program that finds the smallest number in this list
package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	min := x[0] // assume first value is the smallest
	// "i" is 0; if "i" is less than x - after if then i = i + x [i]
	for i := 0; i < len(x); i++ {
		// if it is smaller than the min number
		if x[i] < min {
			min = x[i]
			// found another smaller value, replace previous value in min
		}
	}
	fmt.Println("Lowest number is:", min)
}
