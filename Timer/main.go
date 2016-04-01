package main

import F "fmt"
import "time"

func main() {
    timer1 := time.NewTimer(time.Second * 2)

    <-timer1.C
    F.Println("[*]Timer 1 expired")

    /*
    ** timer2 has to wait for second to execute
    */
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        F.Println("[*]Timer 2 expired")
    }()

    /*
    ** But here we stop timer2, and the goroutine will not to execute
    ** Stop() return bool value
    */
    stop := timer2.Stop()
    if stop {
        F.Println("[*]Timer 2 stoped")
    }

}
