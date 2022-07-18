package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  //Y: 0 is implicit
	v3 = Vertex{}      // X: 0, Y: 0
	p  = &Vertex{1, 2} //has type *Vertex
)

func main() {
	a := Vertex{1, 2}
	b := &a
	b.X = 10

	fmt.Println(a, *b)
}
