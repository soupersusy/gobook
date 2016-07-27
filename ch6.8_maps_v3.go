package main

import "fmt"

func main() {
	//The outer map is used as a lookup table based on the element's symbol,
	//while the inner maps are used to store general information about the
	//elements.
	elements := map[string]map[string]string{
		"H": map[string]string{
			//note commas
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
