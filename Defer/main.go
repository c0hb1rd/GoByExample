package main

import F "fmt"
import "os"

func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

//create a file
func createFile(path string) *os.File {
    defer F.Println("[-]File created.")

    F.Println("[*]File creating...")
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }

    return file
}

//write data to file
func writeFile(file *os.File) {
    defer F.Println("[-]Writed")

    F.Println("[*]Writing...")
    F.Print("[-->]Send data to file:"); var data string; F.Scanln(&data)
    F.Fprintln(file, data)

}

//close file
func closeFile(file *os.File) {
    defer F.Println("[*]File closed.")

    F.Println("[*]Closing file...")
    file.Close()
}
