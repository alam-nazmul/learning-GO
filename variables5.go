package main

import "fmt"

func main() {
	var name string = "KodeKloud"
	var number int = 42
	fmt.Printf("Welcome to %v", name)
	fmt.Printf("Your grade is %d", number)
	fmt.Printf("Welcome to %v! Your grade is %d", name, number)
}
