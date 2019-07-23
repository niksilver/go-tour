package main

import "fmt"

func fibonacci(c chan int, quit chan bool) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <- quit:
            return
        }
    }
}

func main() {
    c := make(chan int)
    q := make(chan bool)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        q <- true
    }()
    fibonacci(c, q)
    fmt.Println("Done")
}

