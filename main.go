// go run main.go gin
// go run main.go naive

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./receipt-processor [naive|gin]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "naive":
		RunNaiveImpl()
	case "gin":
		RunGinImpl()
	default:
		fmt.Println("Invalid argument. Please use 'naive' or 'gin'.")
		os.Exit(1)
	}
}
