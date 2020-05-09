package main
import "fmt"

type Vertex struct {
	X int
	Y int
}

func main(){
	v := Vertex{1, 2}
	p := &v
	v.Y = 4

	fmt.Println(v.Y)
	fmt.Println(p.X)
}