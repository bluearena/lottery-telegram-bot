package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	var i, j int = 1, 2
	var p *int
	p = &i
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	*p = *p - 1
	fmt.Println(i)
	fmt.Println("i,j", i, j)
	fmt.Println("Vertex")
	temp := Vertex{3, 5}
	fmt.Println(temp.X)
}
