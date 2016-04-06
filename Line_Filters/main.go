package main

import(
    "bufio"
    F "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        if ucl == "END" {
            F.Println("Exit.")
            break
        }
        F.Println(ucl)

    }

    if err := scanner.Err(); err != nil {
        F.Println(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
