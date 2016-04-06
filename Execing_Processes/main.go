package main

import(
    "syscall"
    "os"
    "os/exec"
    F "fmt"
)

func main() {
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-alh"}

    env := os.Environ()

    F.Println("$ ls -alh")
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
