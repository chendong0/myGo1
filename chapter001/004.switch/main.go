package main

import "fmt"

func main() {
	var money int = 200
	var busy bool = true
	switch money {
	case 20:
		fmt.Println("点个外卖")
		fallthrough
	case 200:
		fmt.Println("去饭店")
		if busy {
			break
		} else {
			fmt.Println("补充三颗松鼠")
		}
		//todo....

	}
	fmt.Println("end")
}
