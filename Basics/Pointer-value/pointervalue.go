package main

import (
	"fmt"
)

func main() {
	a := 42
	b := a
	c := &a
	fmt.Println("the value of a is", a)
	fmt.Println("the value of b is", b)
	fmt.Println("the value of c is", c)
	a = 50
	fmt.Println("the value of a is", a)
	fmt.Println("the value of b is", b)
	fmt.Println("the value of c is", *c)
	*c = 100
	fmt.Println("the value of a is", a)
	fmt.Println("the value of b is", b)
	fmt.Println("the value of c is", *c)

}
