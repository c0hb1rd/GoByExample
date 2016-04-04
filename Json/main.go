//about this example: http://gobyexample.everyx.in/json/
package main

import (
    "encoding/json"
    "os"
    F "fmt"
)

type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int        `json:"page"`
    Fruits []string `jspm:"fruits"`
}

func main() {
    bolB, _ := json.Marshal(true)
    F.Println(string(bolB))

    intB, _ := json.Marshal(1)
    F.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    F.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    F.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    F.Println(string(slcB))

    mapD := map[string]int{"apple":5, "pear":7}
    mapB, _ := json.Marshal(mapD)
    F.Println(string(mapB))

    res1D := &Response1 {
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    F.Println(string(res1B))

    res2D := &Response2{
        Page: 1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    F.Println(string(res2B))

    byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    F.Println(dat)

    num := dat["num"].(float64)
    F.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    str2 := strs[1].(string)
    F.Println(strs)
    F.Println(str1)
    F.Println(str2)

    str := `{"page": 1, "fruits": ["apple", "pear"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    F.Println(res)
    F.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}

