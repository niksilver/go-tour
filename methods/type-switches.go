package main

import "fmt"

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is \n", v * 2)
    case string:
        fmt.Printf("%v is %d bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know how to deal with something of type %T\n", v)
    }
}

func main() {
    describe(42)
    describe("Hello")
    describe(true)
}

