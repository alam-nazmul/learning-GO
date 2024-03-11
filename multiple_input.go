package main

import "fmt"

func main()  {
	var name string
	var is_muggle bool

	fmt.Print("Enter your name and are you a muggle(T/F): ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println(name, is_muggle)
}