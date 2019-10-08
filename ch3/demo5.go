package main

import "fmt"

func main() {
	s := "Hello world 你好 世界223333"
	fmt.Println("字符串长度:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	s2 := "你好 "
	t := s2
	s2 += "世界!" //字符串可以连接，但原字符串不会改变
	fmt.Println(s2, t)
}
