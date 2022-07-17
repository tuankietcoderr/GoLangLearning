package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	b := &a
	b.X = 10

	fmt.Println(a, *b)
}
