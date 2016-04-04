package main

import S "strings"
import F "fmt"

//The strings package's documents: https://golang.org/pkg/strings/
var P = F.Println
var SOURCE_STRINGS = "I love the computer technology."

func main() {
    P("[*]Source strings:", SOURCE_STRINGS)
    P("[*]Contains \"love\":", S.Contains(SOURCE_STRINGS, "love"))
    P("[*]Count \"o\":", S.Count(SOURCE_STRINGS, "o"))
    P("[*]Has prefix \"I\":", S.HasPrefix(SOURCE_STRINGS, "I"))
    P("[*]Has suffix \"technology.\":", S.HasSuffix(SOURCE_STRINGS, "technology."))
    P("[*]Index of \"p\":", S.Index(SOURCE_STRINGS, "p"))
    P("[*]Join \"And I also love that one\":", S.Join(append(append(make([]string, 0), SOURCE_STRINGS), "I also love that one."), "And "))
    P("[*]Repeat 3 times:", S.Repeat(SOURCE_STRINGS, 3))
    P("[*]Replace all \"t\" to \"T\":", S.Replace(SOURCE_STRINGS, "t", "T", -1))
    P("[*]Replace the first \"t\" character to \"T\":", S.Replace(SOURCE_STRINGS, "t", "T", 1))
    P("[*]Split by blank-space and returns a new slice:", S.Split(SOURCE_STRINGS, " "))
    P("[*]All character to lower:", S.ToLower(SOURCE_STRINGS))
    P("[*]All character to upper:", S.ToUpper(SOURCE_STRINGS))

}
