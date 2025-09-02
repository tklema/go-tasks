package main

import "fmt"

func typeOfVariable(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	// можно так, но как будто так делать плохо
	case chan int, chan string, chan bool, chan any:
		return "cnah"
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
