package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(1000))
	fmt.Println(reflect.TypeOf(10.02))
	fmt.Println(reflect.TypeOf("Uday"))
	fmt.Println(reflect.TypeOf(false))
}
