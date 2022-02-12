package main

import (
	//"fmt"
	"fmt"
	"reflect"
)

func main() {
	var val interface{} = "笑脸" //空接口(interface{})做类型判断,
	//空接口案例 https://segmentfault.com/a/1190000019947900?utm_source=tag-newest
	switch newVal := val.(type) { //type关键字判断,该方法必须适用于switch case中,通过不同的case来进行不同的处理.
	case int:
		tmpVal := newVal + 88.0
		fmt.Println("int", tmpVal)
	case float64:
		tmpVal := newVal + 1.234
		fmt.Println("float64", tmpVal)
	default:
		fmt.Println("money是未知类型")
	}

	retType, val1 := interfaceAssert1(val)
	fmt.Printf("type:%v, value:%v\n", retType, val1)

	retType2, val2 := interfaceAssert2(val)
	fmt.Printf("type:%v, value:%v\n", retType2, val2)

}
func interfaceAssert1(unknown interface{}) (retType, val1 interface{}) {

	//1.直接断言
	val1, ok := unknown.(string)
	if ok {
		return "string", val1
	} else {
		return "not string", nil
	}
}

//3.反射,反射位于relfect包,获取类型使用reflect.TypeOf,获取值使用reflect.ValueOf,具体使用方法:
func interfaceAssert2(unknow interface{}) (retType reflect.Type, val reflect.Value) {
	retType = reflect.TypeOf(unknow)
	return retType, val

}
