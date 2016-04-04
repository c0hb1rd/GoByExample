package main

import "strings"
import F "fmt"

//returns the first index of the target string t, or -1 if no match is found
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }

    return -1
}

//returns true if the target string t is in silce
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}


//return true if one of strings in the slice satisfies the predicate f
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

//returns true if all of the strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return false
        }
    }

    return true
}

//returns a new slice containing all strings in the slice that satify the predicate f
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }

    return vsf
}

//returns a new slice containing the results of applying the function f to each string in the original slice
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }

    return vsm
}

func main() {
    var strs = []string{"peach", "apple", "pear", "watermelon", "strawberry", "orange", "fack you!!!", "Hahaha"}

    F.Println("[*]Source slice:", strs)

    F.Println("[*]Index of \"pear\" in the slice:", Index(strs, "pear"))
    F.Println("[*]It is include \"fack you!!!\":", Include(strs, "fack you!!!"))

    F.Println("[*]Any strings has prefix \"p\":", Any(strs, func (v string) bool { return strings.HasPrefix(v, "p") }))
    F.Println("[*]All strings has prefix \"p\":", All(strs, func (v string) bool { return strings.HasPrefix(v, "p") }))

    F.Println("[*]All strings which contains \"e\" in the slice:", Filter(strs, func(v string) bool {return strings.Contains(v, "e")}))

    F.Println("[*]Upper the first three strings:", Map(strs[0:3], strings.ToUpper))
    F.Println("[*]Upper all strings:", Map(strs, strings.ToUpper))
}
