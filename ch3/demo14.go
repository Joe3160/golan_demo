package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi,i'm Job,Hi."
	fmt.Printf("The postion  of \"Job\" is %d \n", strings.Index(str, "Job"))
	fmt.Printf("The position of first instance of \"Hi\" is %d \n", strings.Index(str, "Hi"))
	fmt.Printf("The position of last instance of \"Hi\" is %d \n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"tom\" is %d \n", strings.Index(str, "tom"))
}
