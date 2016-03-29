package main

import F "fmt"
import "time"

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        F.Println("[*]Channel1 recived.")
        c1 <- "Goodbye"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        F.Println("[*]Channel2 recived.")
        c2 <- "World"
    }()

    F.Println("[*]Working...")

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            F.Println("[->]Channle1 send:", msg1)
        case msg2 := <-c2:
            F.Println("[->]Channel2 send:", msg2)
        }
        time.Sleep(time.Millisecond * 200)
    }
}
