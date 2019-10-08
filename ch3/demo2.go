package main

import "fmt"

func main() {
	var f1, f2 float64
	f1, f1 = 1, 1.0000000000000000001
	if f1 == f2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}
