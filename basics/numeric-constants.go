package main

import "fmt"

const (
    // A really big number
    Big = 1 << 100
    // ...and a smallish one..
    Small =  Big >> 99
)

func makeInt(x int) int {
    return x * 10 + 1
}

func makeFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println("makeInt(Small):", makeInt(Small))
    fmt.Println("makeFloat(Small):", makeFloat(Small))
    fmt.Println("makeFloat(Big):", makeFloat(Big))
}
