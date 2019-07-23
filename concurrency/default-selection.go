package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(1000 * time.Millisecond)
    boom := time.Tick(5000 * time.Millisecond)

    for {
        select {
        case <-tick:
            fmt.Println("Tick")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Print(".")
        }
        time.Sleep(200 * time.Millisecond)
    }
}

