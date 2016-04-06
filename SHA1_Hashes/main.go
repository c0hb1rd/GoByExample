package main

import(
    F "fmt"
    "crypto/sha1"
)

func main() {
    s := "sha1 this string"

    h := sha1.New()

    //if have a strings, use []byte to coerce it to bytes
    h.Write([]byte(s))

    //Sum can be used to append an existing byte slice, it usually is not needed
    bs := h.Sum(nil)

    F.Println(s)
    //SHA1 hash results common printed in hex
    F.Printf("%x\n", bs)
}
