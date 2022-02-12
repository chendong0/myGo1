package main

import "fmt"

func main() {
	hello()
	helloToSomeone("光头强")
	helloToSomeone("熊二")
	msg := constructMessage("熊大")
	fmt.Println(msg)
	calcBMI(1.99, 50.00)
	fmt.Println(calcBMI(1.88,50.00))
	//bmi := calcBMI(1.99, 88.00)
	//fmt.Printf("使用变量的方式调用%.6f", bmi)//正确的小数点格式化符号为%.2f,我写错了%2.f
	//

}

func hello() {
	fmt.Println("hello, golang!")
}

func helloToSomeone(name string) {
	fmt.Println("你好, ", name)

}

func constructMessage(name string) string {
	return "你好, " + name + ", 再来一个"

}

func calcBMI(tall float64, weight float64) float64 { //声明并定义一个返回值为float64数据类型的函数
	return tall / (weight * weight) //如果是多个返回值,用逗号隔开

}
