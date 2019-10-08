package main

import "fmt"

func main() {
	stuNum,_,_:=getClass(20,"一班","张三")
	fmt.Println(stuNum)
}

func getClass(stuNum int, className string, headTeacher string) (int, string, string) {
	return stuNum, className, headTeacher
}
