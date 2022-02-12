package main //使用系统自配的入口,将不兼容报错.

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello,golang!"
	var world = "world"
	fmt.Println(hello, world)
	小数 := 1.234
	fmt.Println(小数)

	var int3, int4 uint = 11, 22
	fmt.Println(int3 * int4)

	ho, ver := 3, 3.123
	var sc = ho * int(ver)
	fmt.Println(ho * int(ver))
	fmt.Println(sc)
	var newname string //只声明,并没有赋值.
	fmt.Println("newname=", newname)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6)
	fmt.Println(int7)

	var nameOfSquare string
	var _name string
	fmt.Println(nameOfSquare, _name)
}
