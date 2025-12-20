package main

import (
	"fmt"
	"strconv"
)

func main() {

	type score struct {
		name  string
		score int
	}

	var scores []score
	scores = []score{}
	i := 0

	for i < 3 {
		fmt.Println("Please enter the student name and score")
		var name, rawscore string
		fmt.Scanln(&name, &rawscore)
		s, _ := strconv.Atoi(rawscore)
		scores = append(scores, score{name: name, score: s})
		i++
	}

	for _, value := range scores {

		fmt.Println("the name of the student is", value.name)
		fmt.Println("the score of the student is", value.score)
	}

}
