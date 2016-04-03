package main

import F "fmt"
import "sort"

func main() {

    F.Println("-------------String Example-------------")
    strs := []string{"c", "a", "b"}
    F.Println("[*]Before sorting:", strs)
    sort.Strings(strs)
    F.Println("[*]After sorted:", strs)


    F.Println("\n------------Integer Example-------------")
    ints := []int{23, 54, 123, 43, 5}
    F.Println("[*]Before sorting:", ints)
    sort.Ints(ints)
    F.Println("[*]After sorted:", ints)

    F.Println("\n------------Weither Sorted--------------")
    s := sort.IntsAreSorted(ints)
    if s {
        F.Println("[*]The integer is sorted.")
    } else {
        F.Println("[*]The integer is not sorted.")
    }
}
