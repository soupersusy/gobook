package main

import "fmt"

func main() {
	// map is an unordered collection of key-value pairs
	// aka assoc array,  hash table or dictionary
	// x is a map of strings to ints
	x := make(map[string]int)
	//  key type in brackets followed by the value type
	// map[string]int
	x["key"] = 10
	fmt.Println(x["key"])
}
