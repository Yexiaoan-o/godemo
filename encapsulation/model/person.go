// Write a program where you cannot access the age and salary fields directly.
// The program also verifies the two fields based on certain rules.
// In the program, the model package calls the main package.

package model
import "fmt"

type person struct {
	Name string
	age int
	salary float64
}

// Define a factory function
func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

// Define methods for age and salary to be accessed outside the package

func (p *person) SetAge(age int) {
	if age > 0 && age < 120 {
		p.age = age
	} else {
		fmt.Println("Invalid age")
	}
}

func (p *person) GetAge(){
	fmt.Println(p.age)
}

func (p *person) SetSalary(salary float64) {
	if salary > 0 {
		p.salary = salary
	} else {
		fmt.Println("Invaid salary")
	}
}

func (p *person) GetSalary(){
	fmt.Println(p.salary)
}