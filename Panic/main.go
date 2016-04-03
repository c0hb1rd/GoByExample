package main

import "os"

func main() {
    /*
    ** here we have a example about panic that is when user creating a file
    */
    panic("[*]A problem, program abort here.")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
