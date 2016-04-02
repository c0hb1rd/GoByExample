package main

import F "fmt"
import (
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    var state = make(map[int]int)
    var mutex = &sync.Mutex{}
    var ops int64 = 0

    //start 100 goroutines to reads
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)

                //ensure only one goroutines could access
                mutex.Lock()
                total += state[key]
                mutex.Unlock()

                atomic.AddInt64(&ops, 1)

                runtime.Gosched()
            }
        }()
    }

    //stat 10 goroutines to simulate writes
    //using the same pattern we did for reads
    for w := 0; w <= 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)

                mutex.Lock()
                state[key] = val
                mutex.Unlock()

                atomic.AddInt64(&ops, 1)

                runtime.Gosched()
            }
        }()
    }

    //wait a second to operations count
    F.Println("[*]Waitting a second to work...")
    time.Sleep(time.Second)

    //report the final count
    opsFinal := atomic.LoadInt64(&ops)
    F.Println("[*]Final ops:", opsFinal)

    //with a final lock of state, show how it ended up
    mutex.Lock()
    F.Println("[*]Final state:", state)
    mutex.Unlock()
}
