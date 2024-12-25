package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "generate":
		generate()
	default:
		fmt.Println("No argument provided")
	}

}
