package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	// var a int
	// a = 1
	// Increase(&a)

	var s Student = Student{name: "John", age: 25, class: "수학", grade: "A+"}

	s.InputSungjuk("과학", "C")
	s.PrintSungjuk()
}

// func Increase(x *int) {
// 	*x = *x + 1
// }
