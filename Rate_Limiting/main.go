package main

import "time"
import F "fmt"

func main() {
    F.Println("[*]Simulation rate-limiting requests:")
    /*
    ** initialization 5 requests
    */
    F.Println("[*]Every 200 millisecond only allow a request")
    requests := make(chan int, 5)
    for i := 1; i < 6; i++ {
        requests <- i
    }
    close(requests)

    //simulation every 200ms only allow deal with a request
    limiter := time.Tick(time.Millisecond * 200)
    for req := range requests {
        <-limiter
        F.Println("[-->]Requests", req, time.Now())
    }

    F.Println("\n[*]Simulation short bursty requests:")


    /*
    ** Simulation short bursty Requests
    */
    burstyLimiter := make(chan time.Time, 3)
    after := make(chan bool)

    //allow 3 request to bursty requests
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    F.Println("[*]Allow 3 requests to short bursty requests")
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    /*
    ** initialization 5 Requests
    */``
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        F.Println("[*]Requests", req, time.Now())
    }
}
