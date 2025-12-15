package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	type student struct {
		name  string
		score int
	}

	s := []student{{name: "maniratnam", score: 90}, {name: "Sireesha", score: 85}}

	fmt.Println("Select score to print option (1-2)")
	var option string
	fmt.Scan(&option)
	index, err := strconv.Atoi(option)

	if err != nil {
		fmt.Println("chosen option is invalid, please choose a correct one")
		return
	}

	switch index {
	case 1:
		index = 0
	case 2:
		index = 1
	default:

	}

	fmt.Println("The students score are as below")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(s[index].name, s[index].score)
}
