package main

import "fmt"

type Person struct {
	Name string
	age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)
}

func main(){
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, ",", z)
}