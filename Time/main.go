package main

import (
    F "fmt"
    "time"
)

func main() {
    now := time.Now()
    F.Print(`[--Func--]now := Time.Now(): `)
    F.Println(now)
    F.Print(`[--Func--]then := Time.Date(2009, 11, 17, 20, 34, 58, 651382737, time.UTC): `)

    then := time.Date(2009, 11, 17, 20, 34, 58, 651382737, time.UTC)
    F.Println(then)

    F.Print(`[--Func--]Year(): `)
    F.Println(then.Year())

    F.Print(`[--Func--]Month(): `)
    F.Println(then.Month())

    F.Print(`[--Func--]Day(): `)
    F.Println(then.Day())

    F.Print(`[--Func--]Hour(): `)
    F.Println(then.Hour())

    F.Print(`[--Func--]Minute(): `)
    F.Println(then.Minute())

    F.Print(`[--Func--]Second(): `)
    F.Println(then.Second())

    F.Print(`[--Func--]Nanosecond(): `)
    F.Println(then.Nanosecond())

    F.Print(`[--Func--]Location(): `)
    F.Println(then.Location())

    F.Print(`[--Func--]Weekday(): `)
    F.Println(then.Weekday())

    F.Print(`[--Func--]Before(now): `)
    F.Println(then.Before(now))

    F.Print(`[--Func--]After(now): `)
    F.Println(then.After(now))

    F.Print(`[--Func--]Equal(now): `)
    F.Println(then.Equal(now))

    diff := now.Sub(then)

    F.Print(`[--Func--]now.Sub(then): `)
    F.Println(diff)

    F.Println()
    F.Print(`[--Func--]diff.Hours(): `)
    F.Println(diff.Hours())

    F.Print(`[--Func--]diff.Minutes(): `)
    F.Println(diff.Minutes())

    F.Print(`[--Func--]diff.Seconds(): `)
    F.Println(diff.Seconds())

    F.Print(`[--Func--]diff.Nanoseconds(): `)
    F.Println(diff.Nanoseconds())

    F.Print(`[--Func--]then.Add(diff): `)
    F.Println(then.Add(diff))

    F.Print(`[--Func--]then.Add(-diff): `)
    F.Println(then.Add(-diff))

}
