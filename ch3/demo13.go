package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "a"))
	fmt.Println(strings.ContainsAny("team", "a b"))
	fmt.Println(strings.ContainsAny("team", "")) //注意：false
	fmt.Println(strings.ContainsAny("", ""))

	fmt.Println()

	fmt.Println(strings.Contains("team", "a"))
	fmt.Println(strings.Contains("team", "a b"))
	fmt.Println(strings.Contains("team", "")) //注意：true
	fmt.Println(strings.Contains("", ""))
}
