package main

import F "fmt"
import "time"

func worker(done chan bool) {
    F.Println("[*]Goroutines working...")
    for i := 5; i > 0; i-- {
        F.Println("[->]Time out:", i)
        time.Sleep(time.Second)
    }

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    if <-done {
        F.Println("[*]Goroutines done.")
    }
}
