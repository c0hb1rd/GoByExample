package main

import F "fmt"
import "time"

func main() {

    c1 := make(chan string, 1)
    c2 := make(chan bool, 1)

    go func() {
        <-c2
        F.Println("[*]Need to waitting 2 second...")
        time.Sleep(time.Second * 2)
        c1 <- "[*]Success recived."
    }()

    var timeout int
    F.Print("[*]Set timeout(between 1 and 3):"); F.Scanln(&timeout);

    c2 <- true

    switch(timeout) {
    case 1:
        select {
        case res := <-c1:
            F.Println(res)
        case <-time.After(time.Second * 1):
            F.Println("[*]Allow only timeout: 1 second.")
        }
    case 2:
        select {
        case res := <-c1:
            F.Println(res)
        case <-time.After(time.Second * 2):
            F.Println("[*]Allow only timeout: 2 second.")
        }
    case 3:
        select {
        case res := <-c1:
            F.Println(res)
        case <-time.After(time.Second * 3):
            F.Println("[*]Allow only timeout: 3 second.")
        }
    default:
        F.Println("[*]Timeout out of allow range.")
    }
}
