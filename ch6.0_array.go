package main

import "fmt"

func main() {
	// array is a numbered sequence of elements of a single type with a fixed length
	// "x" is an array length of 5 float64
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var tt float64 = 0
	//loop to compute average
	// i is the current postion in the array
	for i := 0; i < 5; i++ {
		//below, "i" is the index spot for the array
		// ""+="" means x = x + y or tt=tt+next index float64
		// 0 + 98 then 98 + 93 then 191 + 77...etc..
		tt += x[i]
	}
	fmt.Println(tt / 5)
}
