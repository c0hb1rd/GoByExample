package main

import(
    "os/exec"
    F "fmt"
    "io/ioutil"
)

func main() {
    dateCmd := exec.Command("date")
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }

    F.Println("$ date")
    F.Println(string(dateOut))

    grepCmd:= exec.Command("grep", "grep")

    grepIn, _ :=grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    F.Println("$ grep grep \"hello grep\\ngoodbye grep\"")
    F.Println(string(grepBytes))

    //spawn a full command with a string, we can use bash's -c option
    lsCmd := exec.Command("bash", "-c", "ls -a")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }

    F.Println("$ ls -a")
    F.Println(string(lsOut))

}
