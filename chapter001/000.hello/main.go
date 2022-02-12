package main

import "fmt"

func main() {
	var fruit string = "7apples"
	var wartermallon bool = false //true or false
	if wartermallon {
		fruit = "1 apples "
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy: ", fruit)
}
