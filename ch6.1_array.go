package main

import "fmt"

//goal - get average grade for class
func main() {
	//array - single type, fixed length ex. 5 flt integers or less below
	x := [5]float64{
		//fill each element w/a grade
		98,
		93,
		77,
		82,
		83,
	}
	var tt float64 = 0
	// loop set up for average
	// "i" is 0; if "i" is less? than x; then tt = tt + x [i]
	// tt = 0 + (element 0) 98; tt = 98 + (element 1) 93; etc..
	for i := 0; i < len(x); i++ {
		tt += x[i]
	}
	//divide total by number of elements to get avg.
	fmt.Println(tt / float64(len(x)))
}
