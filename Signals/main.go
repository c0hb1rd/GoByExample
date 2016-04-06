package main

import(
    F "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
        F.Println()
        F.Println(sig)
        done <- true
    }()

    F.Println("[*]Awaiting signal")
    <-done
    F.Println("[*]Exiting.")
}
