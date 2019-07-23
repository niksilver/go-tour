package main

import "fmt"

func sum(s []int, ch chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    ch <- sum
}

func main() {
    s := []int{5, 6, 7, -9, 4, 0}

    ch := make(chan int)
    go sum(s[:len(s)/2], ch)
    go sum(s[len(s)/2:], ch)

    x, y := <-ch, <-ch
    fmt.Println(x, y, x + y)
}

