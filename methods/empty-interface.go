package main

import "fmt"

func main() {
    var i interface{}

    describe(i)

    i = 42
    describe(i)

    i = "Hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%T, %v)\n", i, i)
}

