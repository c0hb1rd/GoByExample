package main

import(
    "bufio"
    F "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data1 := []byte("Hello,world.\n")
    err := ioutil.WriteFile("/tmp/dat1", data1, 0644)

    f, err := os.Create("/tmp/dat2")
    check(err)

    defer f.Close()

    data2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(data2)
    check(err)
    F.Printf("worte %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    F.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    F.Printf("wrote %d bytes\n", n4)

    w.Flush()
}
