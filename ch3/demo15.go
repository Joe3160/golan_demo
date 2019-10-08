package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "花好月圆"
	fmt.Printf("The postion of \"月\" is : %d\n", strings.IndexRune(str, '月'))
}
