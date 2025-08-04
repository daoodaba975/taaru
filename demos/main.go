package main

import "fmt"

// main function
func main() {
	fmt.Println("Hello, world!")
}

// struct
type Person struct {
	Name string
	Age  int
}

// function
func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}
