package main

import(
    F "fmt"
    "time"
)

func main() {
    p := F.Println

    F.Print(`[*]Now: `)
    t := time.Now()
    p(t.Format(time.RFC3339))

    F.Print(`[*]now.Format("3:04PM"): `)
    p(t.Format("3:04PM"))

    //time model must be "Mon Jan _2 15:04:05 2006"
    F.Print(`[*]now.Format("Mon Jan _2 15:04:05 2006"): `)
    p(t.Format("Mon Jan _2 15:04:05 2006"))

    F.Print(`[*]now.Format("2006-01-02T15:04:05.999999-07:00"): `)
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

    F.Print(`[*]time.Parse(time.RFC3339, "2013-11-01T22:08:41+00:00"): `)
    t1, _ := time.Parse(time.RFC3339, "2013-11-01T22:08:41+00:00")
    p(t1)

    form := "3 04PM"
    t2, e := time.Parse(form, "8 41PM")
    F.Print(`[*]time.Parse("3 04PM", "8 41PM"): `)
    p(t2, e)

    F.Print(`[*]fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00"): `)
    F.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    F.Printf(`[*]time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM"): `)
    p(e)
}
