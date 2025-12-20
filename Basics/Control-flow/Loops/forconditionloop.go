package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	type score struct {
		name  string
		score int
	}

	scores := []score{}
	var shouldcontinue bool
	shouldcontinue = true
	for shouldcontinue {

		fmt.Println("1) Enter a score")
		fmt.Println("2) view the report")
		fmt.Println("3) quit")
		fmt.Println()
		fmt.Println("Please select a option")
		var option string
		fmt.Scanln(&option)

		switch option {

		case "1":
			fmt.Println("Enter the student name and score")
			var name, rawscore string
			fmt.Scanln(&name, &rawscore)
			s, _ := strconv.Atoi(rawscore)
			scores = append(scores, score{name: name, score: s})
		case "2":
			fmt.Println("student scores")
			fmt.Println(strings.Repeat("-", 14))
			fmt.Println(scores)
		case "3":
			shouldcontinue = false
		default:
			fmt.Println("Please choose the right option")
		}
	}
}
