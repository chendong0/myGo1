package main

import (
	"fmt"
	"strconv"
)

func main() {
	var c int
	a, b := 18, 9
	f := 511
	g := 49
	fmt.Printf("1 g整型49以2进制方式显示%b\n", g)
	fmt.Printf("2 f整型511以8进制方式显示%o\n", f)
	fmt.Printf("3 a整型18以二进制方式显示%b\n", a)
	fmt.Println(a ^ b)
	c = a ^ b
	//异或^: 二元:a^b 对应位的值不同位1,相同为0
	d := 11011

	//i,err :=strconv.ParseInt("11011",2,27)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	i, _ := strconv.ParseInt("110001", 2, 10)

	fmt.Println("5 _使用二进制码11001转换为整型", i) //十进制的49,用转换为二进制为110001
	//https://c.runoob.com/front-end/58/  //可以使用二进制在线转换工具

	fmt.Printf("6 %d\n", d)
	fmt.Printf("7 c整型9以二进制方式显示%b\n", c)
	fmt.Printf("8 %d\n", c)
	fmt.Println(b ^ a)

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := -2
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item //二元:a ^ b,对应位的值不同为1,相同为0
			/*
				异或^:
					二元:a ^ b,对应位的值不同为1,相同为0
					一元: ^ a
						按位取反:
							1--->0
							0--->1
			*/
		}

	}
	fmt.Println(result)

	//practice :=make([]arr(11,22,33)),练习切片,错误写法
	practice := make([]int, 6, 8) //用make创建切片
	practice1 := []int{11, 22, 33}
	//var practice []type = make([]type,len)
	for i, value := range practice1 {
		fmt.Printf("index%d, value%d \n", i, value)

	}
	fmt.Println(practice)
	fmt.Println(practice1)
}
