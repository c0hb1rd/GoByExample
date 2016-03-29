package main

import F "fmt"

//only can accept a recive status channel
func ping(pings chan<- string, msg string) {
    pings <- msg
}

//the first paremeter is a send status channel, another is a a recive channel
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "Goodbye world")
    pong(pings, pongs)
    F.Println(<-pongs)
}
