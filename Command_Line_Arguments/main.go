package main

import(
    "os"
    F "fmt"
)

func main() {

    F.Println("[*]Arguments:", os.Args)
    F.Println("[*]Argument[0]:", os.Args[0])
    F.Println("[*]Argument[1]:", os.Args[1])
    F.Println("[*]Argument[2]:", os.Args[2])
    F.Println("[*]Argument[3]:", os.Args[3])
}
