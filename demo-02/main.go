package main

//go:generate go run github.com/tmr232/goat/cmd/goater

import (
	"fmt"

	"github.com/tmr232/goat"
)

func Hello(name string) {
	goat.Self().Name("hi")
	fmt.Printf("Hello, %s!\n", name)
}

func Goodbye(name string, formal bool) {
	goat.Self().Name("bye")
	if formal {
		fmt.Printf("Goodbye, %s. Have a good day.\n", name)
	} else {
		fmt.Printf("Bye %s!\n", name)
	}
}

func main() {
	goat.App("greeter",
		goat.Command(Hello),
		goat.Command(Goodbye),
	).Run()
}
