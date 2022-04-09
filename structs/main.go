package main

import "fmt"

// define a struct, which contains the properties and their types
type person struct {
	firstname string
	lastname  string
}

func main() {
	alex := person{firstname: "Alex", lastname: "Anderson"}
	fmt.Println(alex)
}
