package main

//go:generate go run github.com/tmr232/goat/cmd/goater

import (
	"fmt"

	"github.com/tmr232/goat"
)

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	goat.Run(Hello)
}
