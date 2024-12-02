package main

import "fmt"

func initialize() {
	fmt.Println("Initialization complete.")
}

func main() {
	initialize() // Call the renamed function
	fmt.Println("Hello, World!")
}
