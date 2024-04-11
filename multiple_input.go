package main

import "fmt"

<<<<<<< HEAD
func main()  {
	var name string
	var is_muggle bool

	fmt.Print("Enter your name and are you a muggle(T/F): ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println(name, is_muggle)
}
=======
func main() {
	var name string
	var is_muggle bool

	fmt.Printf("Enter your name and your status: ")
	fmt.Scanf("%s, %t", &name, &is_muggle)
	fmt.Println(name, !is_muggle)
}
>>>>>>> 48da66d (multiple input)
