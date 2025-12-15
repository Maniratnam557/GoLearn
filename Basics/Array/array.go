package main

import (
	"fmt"
)

func main() {
	bykes := [3]string{"hero", "honda", "tvs"}
	fmt.Println("List of bykes", bykes)
	fmt.Println("the no of bykes", len(bykes))
	fmt.Println("My Favourite byke is", bykes[1])
}
