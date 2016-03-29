package main

import F "fmt"

func f(from string) {
    for i := 0; i < 10; i++ {
        F.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutines")

    go func(msg string) {
        F.Println(msg)
    }("going")

    var input string
    F.Scanln(&input)
    F.Println("Done")
}
