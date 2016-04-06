package main

import F "fmt"
import "time"

/*
** simulate worker working, worker will stop 2 second after every job done
*/
func Work(id int, jobs <-chan int, result chan<- bool) {
    for j := range jobs {
        F.Println("[-->]Worker", id, "processing job", j, "...")
        time.Sleep(time.Second * 2)
        result <- true
    }
}

func main() {
    jobs := make(chan int, 100)
    result := make(chan bool)

    /*
    ** initialization 3 worker to working
    */
    const WORKER int = 3
    for worker := 1; worker <= WORKER; worker++  {
        go Work(worker, jobs, result)
    }

    /*
    ** initialization 9 jobs
    */
    const JOBS int = 9
    for work := 1; work <= JOBS; work++ {
        jobs <- work
    }
    close(jobs)

    for done := 1; done <= JOBS; done++ {
        if <-result {
            F.Printf("[*]Job %d done\n", done)
        }
    }

    F.Println("[*]All jobs need 18 second to finish, but only spend 6 second to solved it, because we initialization 3 \"worker\" to working.")
}
