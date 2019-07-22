package main

import (
    "fmt"
    "strings"
    //"io"
)

func main() {
    r := strings.NewReader("Hello, worldy!")

    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err != nil {
            fmt.Println(err.Error())
            break
        }
    }
}

