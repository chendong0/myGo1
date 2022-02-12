package main

import "fmt"

func main() {
	var fruit string = "6 walnut"
	var watermelon bool = false //true// true or false
	if watermelon {
		fruit = " 1 banana"
	} else {
		fruit = "8 pear"
	}
	fmt.Println("buy: ", fruit)

}
