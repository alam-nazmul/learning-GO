package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Nazmul"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(100.85))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}