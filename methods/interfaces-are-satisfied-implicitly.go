package main

import "fmt"

type T struct {
    S string
}

type I interface {
    M()
}

func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"Hello, world!"}

    i.M()
}

