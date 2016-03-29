package main

import f "fmt"
import "errors"

func f1(x int) (int, error) {
    if x == 42 {
        return -1, errors.New("Can not work with 42.")
    }

    return  x + 3, nil
}

type argError struct {
    x int
    prob string
}

func (x argError) Error() string {
    return f.Sprintf("%d - %s", x.x, x.prob)
}

func f2(x int) (int, error) {
    if x == 42 {
        return -1, argError{x, "Can not work with it."}
    }
    return x + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            f.Println("Failed", e)
        } else {
            f.Println("Worked", r)
        }
    }

    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            f.Println("Failed", e)
        } else {
            f.Println("Worked", r)
        }
    }
}

