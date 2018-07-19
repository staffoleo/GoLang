package main

import (
	"fmt"
	"os"
)

func main() {
	//Type inference
	args := os.Args

	//Manual type declaratio
	var message string
	message = "Hello, I'm Gopher"

	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println(message)
	}
}
