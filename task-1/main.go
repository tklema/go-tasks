package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
	height  int
	weight  int
}

func (human *Human) getFullName() string {
	return fmt.Sprintf("%s %s", human.name, human.surname)
}

func (human *Human) getFullInfo() string {
	return fmt.Sprintf("name: %s, surname: %s, age: %d, height: %d, weight: %d",
		human.name, human.surname, human.age, human.height, human.weight)
}

type Action struct {
	Human // встраиваем структуру Human (поля и методы)
}

func (action *Action) doAction() {
	fmt.Printf("%s is doing action\n", action.getFullName())
}

func main() {

	action := Action{
		Human: Human{
			name:    "Alex",
			surname: "Soier",
			age:     18,
			height:  180,
			weight:  70,
		},
	}

	fmt.Println(action.getFullName())
	fmt.Println(action.getFullInfo())
	action.doAction()
	fmt.Println(action.age)
}
