package main

import "fmt"

func main() {
	var rmb int = 21

	if rmb <= 20 {
		// 如果有不超过20,点炒粉
		fmt.Println("点个炒粉") // a
	} else if rmb <= 200 {
		//如果不超过200,去餐馆
		fmt.Println("去餐馆") //b
	} else if rmb <= 2000 {
		// 如果不超过2000,去三星级米其林
		fmt.Println("去三星级米其林") //c
	} else if rmb <= 20000 {
		//如果不超过20000,去北京
	} else {
		//如果再多,就去美国
		fmt.Println("定投改变自由")
	}
	fmt.Println("end")
}
