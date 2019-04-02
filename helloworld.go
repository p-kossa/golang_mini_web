package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")

	// var age int
	// or we can do: var age int (defaults to 0)

	var (
		name   = "Kevin"
		age    = 30
		height int
	)

	someVariable := "i am some variable"

	fmt.Println("My name is", name, "age is", age)
	fmt.Println("My height is", height)
	fmt.Println(someVariable)

	fmt.Printf("Type %T value %v", name, name)
}
