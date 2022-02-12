package main

import (
	"fmt"
)

func main() {
	var a, b int8 = 8, 12
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	fmt.Println(true && false == false)
	fmt.Println("a>b=", a > b)
	fmt.Println("a<b=", a < b)

}
