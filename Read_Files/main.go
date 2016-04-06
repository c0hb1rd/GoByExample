package main

import(
    "bufio"
    F "fmt"
    "io"
    "io/ioutil"
    "os"
)

//because reading files requires checking most calls for errors
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    //basic file reading
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    F.Println(`[--Func--]ioutil.ReadFile("/tmp/dat"):`)
    F.Println(string(dat))

    //obtain an os.File value then we can use it do some operation we need
    f, err := os.Open("/tmp/dat")
    check(err)
    //set a buffer to save some datas
    b1 := make([]byte, 5)
    //return how many bytes to read and save datas to b1
    n1, err := f.Read(b1)
    check(err)
    F.Printf("[--Func--]Get os.File value via os.Open(\"/tmp/dat\").\n")
    F.Printf("[--Oper--]Set a buffer([]byte, 5) to save some datas.\n")
    F.Printf("[--Func--]Read(os.File):\n")
    F.Printf("[-->]Bytes: %d\n[-->]Datas: %s\n", n1, string(b1))

    F.Println()

    //seek to a known locatin in the file and read from there
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 4)
    n2, err := f.Read(b2)
    check(err)
    F.Printf("[--Func--]Get os.File value via os.Open(\"/tmp/dat\").\n")
    F.Printf("[--Func--]Seek(6, 0).\n")
    F.Printf("[--Oper--]Set a buffer([]byte, 4) to save some datas.\n")
    F.Printf("[--Func--]Read(os.File):\n")
    F.Printf("[-->]Bytes: %d\n[-->]Seek: %d\n[-->]Data: %s\n", n2, o2, string(b2))

    F.Println()

    //example for io.ReadAtLeast
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 10)
    n3, err := io.ReadAtLeast(f, b3, 10)
    check(err)
    F.Printf("[--Func--]Seek(6, 0).\n")
    F.Printf("[--Oper--]Set a buffer([]byte, 10) to save some datas.\n")
    F.Printf("[--Func--]io.ReadAtLeast(os.File, buffer, 10):\n")
    F.Printf("[-->]Bytes: %d\n[-->]Seek: %d\n[-->]Data:\n%s\n", n3, o3, string(b3))

    //there is no built-in rewind, but we can use Seek(0, 0) to accpmplishes this
    _, err = f.Seek(0, 0)
    check(err)

    F.Println()

    //bufio buffered reader
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    F.Printf("[--Func--]buifo.NewReader(os.File).\n")
    F.Printf("[--Func--]reader.Peek(5).\n")
    F.Printf("[-->]Bytes: %d\n", 5)
    F.Printf("[-->]Data: %s\n", string(b4))
}


