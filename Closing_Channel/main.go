package main

import F "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            //special 2-value, fitst one is integer, next one is boolean
            j, more := <-jobs
            if more {
                F.Println("[*]Recived job", j)
            } else {
                F.Println("[*]Recived all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 0; j < 5; j++ {
        jobs <- j+1
    }
    close(jobs)

    if <-done {
        F.Println("[*]Done.")
    }


}
