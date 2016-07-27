package main

import "fmt"

func main() {
	//[5] elements
	x := [5]float64{
		98,
		93,
		77,
		82,
		// can easily take out a test or have any # of scores less that 5
		//    83,
	}
	var tt float64 = 0
	// _ (underscore) is used to tell the compiler that we don't need this.
	//In this case we don't need the iterator variable for this loop
	for _, value := range x {
		// tt = 0 + 98 (first element in array); loop tt = 98 + 93 (2nd element in array) etc..
		tt += value
	}
	// len(x) so it automatically knows and not fixed
	fmt.Println(tt / float64(len(x)))
}
