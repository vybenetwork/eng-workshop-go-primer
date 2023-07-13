package main

import (
	"context"
	"eng-workshop-go-primer/examples/person"
	"fmt"
	"os"
)

// Declaring custom types from primitives
type Number int

// Not exposed to other packages
func sum(x, y int) Number {
	result := x + y // result is an int

	return Number(result) // type conversion
}

// Exposed to other packages
func Multiply(x, y int) int {
	return x * y
}

const X = 2 // exposed type inferrence
const y = 2 // unexposed

func main() {
	// Variable declaration
	var a string

	// Variable assignment
	a = "a"
	fmt.Printf("var a = %s\n", a)

	// Varible definition shorthand
	sum := sum(X, y)
	fmt.Printf("Sum %d + %d = %d\n", X, y, sum)

	// Standard practive for using single character var names
	// Importing another package. Shows that the internal property `email` is not printed.
	p := person.NewPerson("Matt", "G", "matthew@vybenetwork.com")
	// email := p.email
	fmt.Printf("Email not accessible: %v\n", p)

	// Maps
	dict := map[string]int{
		"One":   1,
		"Twp":   2,
		"Three": 3,
	}

	fmt.Printf("dict.One: %d\n", dict["One"])
	fmt.Printf("dict.Two: %d\n", dict["Two"])
	fmt.Printf("dict.Three: %d\n", dict["Three"])

	// Array
	list := [3]int{1, 2, 3}
	fmt.Printf("%v\n", list)

	// Slice
	slice := make([]string, 0)
	slice = append(slice, "one")
	fmt.Printf("%v\n", slice)

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Create a new channel that receives Number types
	messageChannel := make(chan Number)

	// Go routine - async processing
	go IncrementalSum(ctx, 0, 10, messageChannel)

	for {
		select {
		case count := <-messageChannel:
			fmt.Printf("%d\n", count)

			if count == 5 {
				cancel()
			}
		case <-ctx.Done():
			fmt.Println("That's it for examples.")
			os.Exit(0)
		}
	}
}
