package main

import "fmt"

func main() {
	a := 20
	c := 200
	c = a
	fmt.Println("赋值操作，把 a 赋值给 c，所以 c 的值为：", c)
	c += a
	fmt.Println("相加和赋佳运算符，实际为 c = c + a ，所以 c 的值为： ", c)
	c -= a
	fmt.Println("相减和赋佳运算符，实际为 c = c - a ，所以 c 的值为：", c)
	c *= a
	fmt.Println("相来和赋佳运算符，实际为 c = c * a ， 所以 c 的值为:", c)
	c /= a
	fmt.Println("相除和赋值运算符，实际为 c = c I a ，所以 c 的值为：", c)
	c <<= 2
	fmt.Println("左移和赋佳运算符，所以 c 的值为 :", c)
	c >>= 2
	fmt.Println("右移和赋佳运算符，所以 c 的值为 :", c)
	c &= 2
	fmt.Println("按位与和赋佳运算符，所以 c 的值为： ", c)
	c ^= 2
	fmt.Println("按住异或和赋佳运算符，所以 E 的值为 ：", c)
	c |= 2
	fmt.Println("按位或和赋佳运算符，所以 c 的值为：", c)
}
