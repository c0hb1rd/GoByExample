package main

import F "fmt"
import (
    "time"
    "sync/atomic"
    "runtime"
)

func main() {
    var ops uint64 = 0

    //start 50 gorountine
    for i := 0; i< 50; i++ {
        go func() {
            for {
                //use AddUint64 to auto accmulating ops
                atomic.AddUint64(&ops, 1)

                //allow other goroutines to proceed
                runtime.Gosched()
            }
        }()
    }

    //wait a second to allow some ops to accumulate
    F.Println("[*]Wait 2 second to accumlating...")
    time.Sleep(time.Second * 2)

    opsFinal := atomic.LoadUint64(&ops)
    F.Println("[*]Finall ops:", opsFinal)
}
