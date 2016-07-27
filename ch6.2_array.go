package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}
	// prefer to have the #'s on their own line to comment out
	// more flexible
	var tt float64 = 0
	for i := 0; i < len(x); i++ {
		tt += x[i]
	}
	fmt.Println(tt / float64(len(x)))
}
