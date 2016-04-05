package main

import(
    F "fmt"
    "time"
)

func main() {

    F.Print(`[--Func--]now := time.Now(): `)
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    F.Println(now)

    F.Print(`[--Func--]now.Unix(): `)
    F.Println(secs)

    F.Print(`[--Func--]now.UnixNano(): `)
    F.Println(nanos)

    F.Print(`[--Func--]now.UnixNano() / 1000000: `)
    millis := nanos / 1000000
    F.Println(millis)

    F.Print(`[--Func--]time.Unix(secs, 0): `)
    F.Println(time.Unix(secs, 0))

    F.Print(`[--Func--]time.Unix(0, nanos): `)
    F.Println(time.Unix(0, nanos))
}
