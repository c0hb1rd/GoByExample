package main

import F "fmt"

type geometry interface {
    Area() float64
    Perm() float64
}

type rect struct {
    width float64
    height float64
}

func (x rect) Area() float64 {
    return x.width * x.height
}

func (x rect) Perm() float64 {
    return 2*x.width + 2*x.height
}

func measure(g geometry) {
    F.Println(g.Area())
    F.Println(g.Perm())
    F.Println(g)
}

func main() {
    c := rect{width: 2, height: 3}
    measure(c)
}
