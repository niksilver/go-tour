package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (aged %v)", p.Name, p.Age)
}

func main() {
    ad := Person{"Arthur Dent", 42}
    zb := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(ad)
    fmt.Println(zb)
}
