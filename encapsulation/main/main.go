package main

import (
	"fmt"
	"encapsulation/model"
)



func main() {
	person := model.NewPerson("Sean")

	fmt.Println(person)

	// fmt.Println(person.age)  // person.age undefined, which cannot be accessed directly
	// fmt.Println(person.salary)  // same as person.age

	person.SetAge(10)
	person.GetAge()

	person.SetSalary(10000)
	person.GetAge()

	fmt.Println(person)

}

