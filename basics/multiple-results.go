package main

import "fmt"

func swap(a, b string) (string, string) {
    return b, a
}

func main() {
    i, j := swap("hello", "world")
    fmt.Println(i, j)
}

