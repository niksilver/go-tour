package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i += 1 {
        z -= (z*z - x) / (2 * z)
    }
    return z
}

func Compare(x int) {
    ours := Sqrt(float64(x))
    official := math.Sqrt(float64(x))
    fmt.Println(
        "Input:", x,
        "Ours:", ours,
        "Official:", official,
        "Difference:", (ours - official),
    )
}

func main() {
    for i := 0; i < 200; i += 10 {
        Compare(i)
    }
}

