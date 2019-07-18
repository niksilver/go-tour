package main

import "fmt"

// On successive calls what's returned from
// this function return the next
// Fibonacci number in the sequence.
// Starts off 0, 1, 1, 2.

func fibonacci() func() int {
    prev1, prev2, counter := 0, 0, 0
    return func() int {
        var next int
        switch counter {
            case 0:
                next = 0
            case 1:
                prev2, prev1, next = 0, 1, 1
            default:
                next = prev2 + prev1
                prev2 = prev1
                prev1 = next
        }
        counter++
        return next
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
