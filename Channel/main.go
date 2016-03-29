package main

import F "fmt"
import "time"

func main() {
    message := make(chan string)

    F.Println("[*]Message start listening...")
    go func() { F.Println(<-message); F.Println("[*]Recived message.") }()

    var input string
    F.Print("[->]Send a value to message:")
    F.Scanln(&input); message <- input

    time.Sleep(time.Millisecond)
}
