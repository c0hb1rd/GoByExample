package main

import(
    "time"
    F "fmt"
    "math/rand"
)

func main() {

    //returns a random integer n, 0 <= n <= 100
    F.Println(`[*]rand.Intn(100):`, rand.Intn(100))

    //returns a random float64 n
    F.Println(`[*]rand.Float64():`, rand.Float64())

    F.Println(`[*]rand.Float64()*5+5:`, (rand.Float64()*5) + 5)
    F.Println(`[*]rand.Float64() * 5 + 5:`, (rand.Float64() * 5) + 5)

    sd := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(sd)

    F.Printf("\n[*]-----Set a new srand by \"time.Now().UnixNano\" as r1-----\n")
    F.Println(`[*]r1.Intn(100):`, r1.Intn(100))

    F.Printf("\n[!!]-----If seed a source with the same number, it produces the same sequence of random numbers-----\n")
}
