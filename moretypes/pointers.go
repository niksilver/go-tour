package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println("What's at the end of p?", *p)
    *p = 43
    fmt.Println("What's in i?", i)

    p = &j
    *p = *p / 37
    fmt.Println("What's in j?", j)
    j = *p + 1
    fmt.Println("What's at the end of p?", *p)

    fmt.Println("And what is p, exactly?", p)
}

