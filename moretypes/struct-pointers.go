package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{3, 7}
    p := &v
    p.X = 1e9
    fmt.Println("Our vertext is now", v)
}

