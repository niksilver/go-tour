package main

import "fmt"

func main() {
    var i interface{} = "Hello"

    s1 := i.(string)
    fmt.Println(s1)

    s2, ok := i.(string)
    fmt.Println(s2, ok)

    f1, ok := i.(float64)
    fmt.Println(f1, ok)

    f2 := i.(float64)  // Panic
    fmt.Println(f2)
}
