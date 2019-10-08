package main

import "fmt"

func main() {
	s := "我"
	t := "们"
	if true {
		fmt.Println(s[0], t[0])
		fmt.Println(s[1], t[1])
		fmt.Println(s[2], t[2])
	}

	a := "a"
	b := "b"
	if a < b {
		fmt.Println(a[0], "小于", b[0])
	}
}
