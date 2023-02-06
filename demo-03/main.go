package main

import (
	"fmt"

	"github.com/tmr232/goat"
)

//go:generate go run github.com/tmr232/goat/cmd/goater

// Greet greets the given name.
func Greet(name string, goodbye bool, times int, extra *string) {
	goat.Flag(name).Usage("The name to greet")
	goat.Flag(goodbye).Usage("Say goodbye instead of hello")
	goat.Flag(times).Usage("Repeat greeting multiple times").Default(1)
	goat.Flag(extra).Usage("Optionally add an extra `message`")

	for i := 0; i < times; i++ {
		if goodbye {
			fmt.Printf("Goodbye, %s.\n", name)
		} else {
			fmt.Printf("Hello, %s!\n", name)
		}
		if extra != nil {
			fmt.Println(*extra)
		}
	}
}

func main() {
	goat.Run(Greet)
}
