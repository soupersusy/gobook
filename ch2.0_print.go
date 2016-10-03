package main

import "fmt"

var p string = "cat"

func main() {
	var x string
	x = "Hello World TEST"
	fmt.Println(x)
	x = x + " second"
	fmt.Println(x)
	var y string
	y = "From Susy"
	fmt.Println(x == y)
	var z string
	z = "From Susy"
	fmt.Println(y == z)
	m := "short"
	fmt.Println(m)
	q := 5
	fmt.Println(q + q)
	h := "dog"
	fmt.Println(h)
	dogsName := "Max"
	fmt.Println("My ", h+"s name is", dogsName)
	fmt.Println(p)
}
