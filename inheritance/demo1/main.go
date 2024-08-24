// Define a struct that includes common fields and methods
// Define instances that inherit the struct

package main
import "fmt"

// Define a parent struct
type Student struct {
	Name string
	Age int
	Score float64
}

// Define methods for the parent struct
func (stu *Student) ShowInfo(){
	fmt.Printf("Student name is %v, Student age is %v, Student score is %v", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64){
	stu.Score = score
}

// Define a child struct (pupil)
type Pupil struct {
	Student
}

func (p *Pupil) ShowStatus(){
	fmt.Println("The pupil is doing an exam...")
}

// Define a child struct (graduate)
type Graduate struct {
	Student
}

func (p *Graduate) ShowStatus(){
	fmt.Println("The graduate is doing an exam...")
}



func main(){
	// Define an pupil instance
	pupil := Pupil{}
	pupil.Name = "Sean"
	pupil.Age = 10

	pupil.SetScore(100)

	pupil.ShowInfo()
	fmt.Println()
	pupil.ShowStatus()

}


