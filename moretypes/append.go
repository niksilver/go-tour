package main

import "fmt"

func main() {
    var s []int
    printSlice(s)

    // Append to a nil slice
    s = append(s, 0)
    printSlice(s)

    // Append another
    s = append(s, 1)
    printSlice(s)

    // Append several
    s = append(s, 2, 3, 4, 5)
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}

