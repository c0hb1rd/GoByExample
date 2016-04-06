package main

import(
    "os"
    F "fmt"
    "strings"
)

func main() {
    os.Setenv("MYTEST", "1")
    F.Println("MYTEST:", os.Getenv("MYTEST"))
    F.Println("BAR:", os.Getenv("BAR"))

    F.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        F.Printf("[*]%s: %s\n", pair[0], pair[1])
    }
}
