package main

import (
    "fmt"
    "math"
)

type Vector struct {
    X, Y float64
}

func (v Vector) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    vec := Vector{3, 4}
    fmt.Println(vec.Abs())
}

