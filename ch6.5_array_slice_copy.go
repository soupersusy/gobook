package main

import "fmt"

func main() {
	sliceEx1 := []float64{1, 2, 3}
	//make
	sliceEx2 := make([]float64, 2)
	//copy - sliceEX1 is copied in to sliceEx2 but only rm for 2 elements
	copy(sliceEx2, sliceEx1)
	fmt.Println(sliceEx1, sliceEx2)
	//tes
}
