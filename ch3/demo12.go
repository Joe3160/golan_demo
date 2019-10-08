package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is a example of string"
	fmt.Printf("Does the string \"%s\" have prefix \"%s\"?\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("Does the string \"%s\" have subfix \"%s\"?\n", str, "string")
	fmt.Printf("%t\n", strings.HasSuffix(str, "string"))
	fmt.Printf("Does the string \"%s\" content \"%s\"?\n", str, "example")
	fmt.Printf("%t\n", strings.Contains(str, "example"))
}
