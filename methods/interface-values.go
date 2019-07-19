package main

import (
    "fmt"
    "math"
)

type T struct {
    S string
}

type I interface {
    M()
}

type F float64

func (t *T) M() {
    fmt.Println(t.S)
}

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    i = &T{"Hello, world!"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

