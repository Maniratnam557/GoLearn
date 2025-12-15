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

	fmt.Println("The students score are as Below:")
	fmt.Println(strings.Repeat("-", 14))

	if err == nil && index == 1 {
		fmt.Println(s[index-1].name, s[index-1].score)
	} else if err == nil && index == 2 {
		fmt.Println(s[index-1].name, s[index-1].score)
	} else {
		fmt.Println("chossed option is invalid , Please choose correct one")
	}

}
