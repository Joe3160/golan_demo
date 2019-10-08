package main

import "fmt"

func main() {
	v2 := 3.2 + 12i
	v3 := complex(3.2, 12)
	v := v2 + v3
	fmt.Println(v2, v3, v)

	vr := real(v) //获取实部
	vi := imag(v)  //获取虚部
	fmt.Println(vr, vi)
}
