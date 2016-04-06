package main

import(
    "flag"
    F "fmt"
)

func main() {
    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 43, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    flag.Parse()

    F.Println("[*]Support arguments: word, numb, fork, svar, tail")
    F.Println("word:", *wordPtr)
    F.Println("numb:", *numbPtr)
    F.Println("fork:", *boolPtr)
    F.Println("svar:", svar)
    F.Println("tail:", flag.Args())
}
