package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{3, 5}
    v.X = 4
    fmt.Println(v.X)
    fmt.Println(v.Y)
}
