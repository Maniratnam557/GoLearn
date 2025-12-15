package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	students := []string{"Sireesha", "Maniratnam",
		"Prakash", "kalai",
	}
	scores := []int{90, 89, 85, 75}

	fmt.Println("Students Score")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(students[0], scores[0])

	students = append(students, "Divya", "Deepak")

	fmt.Println("Students Score")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("No:of students", len(students))
	fmt.Println(students[0], scores[0])
	fmt.Println(students)

	students = slices.Delete(students, 1, 3)
	fmt.Println(students)

}
