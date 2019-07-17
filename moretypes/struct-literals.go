package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X and Y are 0 implicitly
    p = &Vertex{3, 4}  // has type *Vertex
)

func main() {
    fmt.Printf("v1 %T %v\n", v1, v1)
    fmt.Printf("v2 %T %v\n", v2, v2)
    fmt.Printf("v3 %T %v\n", v3, v3)
    fmt.Printf("*p %T %v\n", *p, *p)
    fmt.Printf("p %T %v\n", p, p)
}

