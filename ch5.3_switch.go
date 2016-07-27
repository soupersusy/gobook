package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		//"switch" followed by "case" short way to do "else if"
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
