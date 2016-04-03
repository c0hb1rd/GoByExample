package main

import F "fmt"
import "sort"

//only alias []string
type ByLength []string

/*
** in order to use sort package's genric sort function
** we need to implement sort.Interface-Len, Less, Swap
** the more important one is Less, if define the sorting logic
*/

//return lenght of slice
func (s ByLength) Len() int {
    return len(s)
}

//swap the element of slice
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

//sorting logic
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits:= []string{"banana", "apple", "orange", "waterm", "strawberry", "watermelon"}

    F.Println("[*]Before sorting:", fruits)
    sort.Sort(ByLength(fruits))
    F.Println("[*]After sorted:", fruits)

}
