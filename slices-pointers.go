package main

import "fmt"

func main(){
	names := [4]string{
		"tamer",
		"erin",
		"ishta",
		"",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:4]
	fmt.Println(a,b)

	b[2] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	ab := []int{2,3,5,7,11,13}
	abd := ab[1:4]
	fmt.Println(abd)

	abc := ab[:2]
	fmt.Println(abc)

	abe := ab[1:]
	fmt.Println(abe)
}