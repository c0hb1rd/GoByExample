package main

import F "fmt"

func main() {
    message := make(chan string)
    signals := make(chan string)

    select {
    case msg := <-message:
        F.Println("[*]Recived message from \"message\":", msg + ".")
    default:
        F.Println("[*]No message recived.")
    }

    msg := "Goodbye World"
    select {
    case message <- msg:
        F.Println("[->]Send message to \"message\":", msg + ".")
    default:
        F.Println("[*]No message recived.")
    }

    select {
    case msg := <-message:
        F.Println("[*]Recived message from \"message\":", msg + ".")
    case sig := <-signals:
        F.Println("[*]Revived message from \"signals\"", sig + ".")
    default:
        F.Println("[*]No message recived.")
    }
}
