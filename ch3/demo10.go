package main

import "fmt"

func main() {
	s:="Hello 世界! "
	b:=[]byte(s)
	b[5]=','
	fmt.Printf("%s\n",s)
	fmt.Printf("%s\n",b)
}
