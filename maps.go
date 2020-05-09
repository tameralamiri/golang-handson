package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var l = map[string]Vertex{
	"Bell Labs": { 40.68433, -74.39967},
	"Google": { 37.42202, -122.08408},
}

func main(){
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("the value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

}