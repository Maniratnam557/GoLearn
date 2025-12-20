package main

import (
	"fmt"
	"strconv"
	"strings"
)

type score struct {
	name  string
	score int
}

func main() {

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
			scores = append(scores, addStudent())
		case "2":
			printReport(scores...)
		case "3":
			shouldcontinue = false
		default:
			fmt.Println("Please choose the right option")
		}
	}
}

func printReport(scores ...score) { /* variadic parameter other way we can use func printReport(scores []score) */
	fmt.Println("student scores")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(scores)
}

func addStudent() score {
	fmt.Println("Enter the student name and score")
	var name, rawscore string
	fmt.Scanln(&name, &rawscore)
	s, _ := strconv.Atoi(rawscore)
	return score{name: name, score: s}

}
