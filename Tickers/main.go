package main

import F "fmt"
import "time"

func main() {
    ticker := time.NewTicker(time.Millisecond * 500)

    /*
    ** this goroutine will always work until we stop it
    ** to execute every 5ms
    */
    go func() {
        for t := range ticker.C {
            F.Println("[*]Ticker at", t)
        }
    }()

    /*
    ** delay 23ms and stop it
    */
    time.Sleep(time.Millisecond * 2300)
    ticker.Stop()
    F.Println("[*]Ticker stoped")

}
