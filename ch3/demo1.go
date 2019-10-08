package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value1 float64
	value1 = 1    //不管是否有小数点 ,value1 都是浮点型
	value2 := 2   //如采没有小数点， value2 会被推导为整型
	value3 := 3.0 //如采没有小数点， value3 会被椎导为i手点型
	//
	//v := value1 + value2 //编译失败， value1 与 value2 类型不同
	v := value1 + value3
	fmt.Println(value1, value2, value3, v)
	fmt.Println("v的类型:", reflect.TypeOf(v))
}
