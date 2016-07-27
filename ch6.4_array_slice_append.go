package main

import "fmt"

func main() {
	// slice is a segment of an array (length adjustable)
	sliceVar1 := []int{1, 2, 3}
	sliceVariableEx2 := append(sliceVar1, 4, 5)
	//  "append" creates a new slice using sliceEx1 with 4,5 to make sliceEx2
	fmt.Println(sliceVar1, sliceVariableEx2)
}
