package main

import(
    "strconv"
    F "fmt"
)

func main() {
    f, _ := strconv.ParseFloat("1.234", 64)
    F.Println(`[*]strconv.ParseFloat("1.234", 64):`, f)

    i, _ := strconv.ParseInt("123", 0, 64)
    F.Println(`[*]strconv.ParseInt("123", 0, 64):`, i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    F.Println(`[*]strconv.ParseInt("0x1c8", 0, 64):`, d)

    u, _ := strconv.ParseUint("789", 0, 64)
    F.Println(`[*]strconv.ParseUint("178", 0, 64):`, u)

    k, _ := strconv.Atoi("135")
    F.Println(`[*]strconv.Atoi("135"):`, k)

    _, e := strconv.Atoi("wat")
    F.Println(`[*]strconv.Atoi("wat"):`, e)

}
