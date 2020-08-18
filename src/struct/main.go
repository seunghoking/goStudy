package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {
	p := Person{name: "개똥이", age: 15}

	fmt.Println(p)

	p.name = "Smith"
	p.age = 24

	fmt.Println(p)

	p.PrintName()
}
