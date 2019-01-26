package main

import (
	"fmt"
	"os"

	"github.com/hb-go/wire/sample"
)

func main() {
	e, err := sample.InitializeEvent("Hello")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
