package main

import (
    "bytes"
    "regexp"
)
import F "fmt"

func main() {
    r, _ := regexp.Compile("p([a-z]+)ch")
    F.Println("[*]---------- Source Compile(\"p([a-z]+)ch\") ----------")
    F.Println("[--Func--]MatchString(\"peach\"):", r.MatchString("peach"))
    F.Println("[--Func--]FindStringIndex(\"peach punch\"):", r.FindStringIndex("peach punch"))
    F.Println("[--Func--]FindStringSubmatch(\"peach punch\"):", r.FindStringSubmatch("peach punch"))
    F.Println("[--Func--]FindStringSubmatchIndex(\"peach punch\"):", r.FindStringSubmatchIndex("peach punch"))
    F.Println("[--Func--]FindAllString(\"peach punch pinch\", -1):", r.FindAllString("peach punch pinch", -1))
    F.Println("[--Func--]FindAllStringSubmatchIndex(\"peach punch pinch\", -1):", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    F.Println("[--Func--]Match([]byte(\"peach\")):", r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    F.Println("\n[*]---------- Source MustCompile(\"p([a-z]+)ch\") ----------")
    F.Println("[--Func--]ReplaceAllstring(\"peach\", \"<fruit>\"):", r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    F.Println("[--Func--]ReplaceAllFun([]byte(\"a peach\"), bytes.ToUpper):", out)
}
