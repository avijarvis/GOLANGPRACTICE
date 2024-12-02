package main

import "fmt"

func main() {
	person := make(map[string]string)
	person["name"] = "John"
	person["age"] = "30"

	fmt.Println(person)         // Output: map[age:30 name:John]
	fmt.Println(person["name"]) // Output: John
}
