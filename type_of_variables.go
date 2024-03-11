package main

import "fmt"

func main()  {
	var grades int = 30
	var name string = "Nazmul"
	var isCheck bool = true
	var amount float32 = 5432.85

	fmt.Printf("variable grades = %v is a type of %T \n", grades, grades)
	fmt.Printf("variable name = %v is a type of %T \n", name, name)
	fmt.Printf("variable isCheck = %v is a type of %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is a type of %T \n", amount, amount)
}