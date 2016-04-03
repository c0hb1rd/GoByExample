package main

import F "fmt"
import (
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {
    var ops int64


    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()


    //start 100 goroutines to reads
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key: rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    //start 100 goroutines to writes
    for w := 0; w < 100; w++ {
        go func() {
            for {
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(100),
                    resp: make(chan bool),
                }
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    //work 1 second
    F.Println("[*]Waitting a second to working...")
    time.Sleep(time.Second)

    //report the final value of operation counter
    F.Println("[*]final value of operation counter:", atomic.LoadInt64(&ops))


}
