package main

import (
	"fmt"
	"reflect"
)

func typeOfVariable(variable interface{}) string {
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "unknown"
	}

}

func main() {
	var intVar int
	var stringVar string
	var boolVar bool
	var chanIntVar chan int
	var chanStringVar chan string
	var chanAnyVar chan any

	fmt.Printf("intVar type: %s\n", typeOfVariable(intVar))
	fmt.Printf("stringVar type: %s\n", typeOfVariable(stringVar))
	fmt.Printf("boolVar type: %s\n", typeOfVariable(boolVar))
	fmt.Printf("chanIntVar type: %s\n", typeOfVariable(chanIntVar))
	fmt.Printf("chanStringVar type: %s\n", typeOfVariable(chanStringVar))
	fmt.Printf("chanAnyVar type: %s\n", typeOfVariable(chanAnyVar))
}
