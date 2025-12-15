package main

import (
	"fmt"
	"strings"
)

func main() {

	name, score := "Maniratnam", 90
	pass := true

	fmt.Println("Student Scores Pass/Fail")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(name, score, pass)
}
