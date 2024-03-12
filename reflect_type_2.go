package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var name string = "Nazmul"
	var id_no int	= 10

	fmt.Printf("Your name is='%v' and it's type is %v \n", name, reflect.TypeOf(name))
	fmt.Printf("Your id is %v and it's type is %v \n", id_no, reflect.TypeOf(id_no))
}