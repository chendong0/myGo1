package main

import "fmt"

func main() {
	var money interface{} = 10
	//var money interface {}=10.0
	//var money interface {}="10"
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3.0
		fmt.Println("money 是 int ", tmpMoney)
	}

}
