package main

import (
	"fmt"
	"reflect"
)

func typeOfVariable(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		// можно так, но как будто тогда зачем 'variable.(type)', если
		// можно сразу сделать с помощью reflect
		if reflect.TypeOf(variable).Kind() == reflect.Chan {
			return "chan"
		}
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
