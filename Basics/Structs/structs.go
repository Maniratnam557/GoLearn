package main

import (
	"fmt"
	"strings"
)

func main() {

	type student struct {
		name  string
		score int
	}

	s := []student{{name: "maniratnam", score: 90}, {name: "Sireesha", score: 85}}

	fmt.Println("The students score are as below")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(s[0].name, s[0].score)
	fmt.Println(s[1].name, s[1].score)
}
