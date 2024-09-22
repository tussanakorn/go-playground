package main

import (
	"fmt"

	"github.com/google/uuid"
	sample "github.com/tussanakorn/go-playground/project"
)

func main() {
	fmt.Println("Hello, World!")
	sample.HelloSample()

	// Generate a UUID
	id := uuid.New()
	fmt.Println("UUID:", id)
}
