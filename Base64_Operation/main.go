package main

import(
    b64 "encoding/base64"
    F "fmt"
)

func main() {

    data := "I love computer technology."
    /*
    ** standard base64 encoder/decoder
    ** requires a []byte, so we cast our string to that type
    */
    dataEncoded := b64.StdEncoding.EncodeToString([]byte(data))
    dataDecoded, _ := b64.StdEncoding.DecodeString(dataEncoded)

    F.Println(`[*]Source data:`, data)
    F.Println(`[*]After base64 encoding:`, string(dataEncoded))
    F.Println(`[*]After base64 decoding:`, string(dataDecoded))

    F.Println()

    urlData := "This is url datas."
    /*
    ** URL-compatible base64 encoder/decoder
    */
    urlEncoded := b64.URLEncoding.EncodeToString([]byte(urlData))
    urlDecoded, _ := b64.URLEncoding.DecodeString(urlEncoded)
    F.Println(`[*]Source url data:`, urlData)
    F.Println(`[*]After URL-compatible encoded:`, string(urlEncoded))
    F.Println(`[*]After URL-compatible decoded:`, string(urlDecoded))
}
