package main

import "fmt"

func main() {
	s := "我是中国人"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])//输出羊字节值
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%d:%c\n", i, v)
	}
}
