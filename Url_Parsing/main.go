package main

import(
    F "fmt"
    "net"
    "net/url"
)

func main() {
    sUrl := "postgres://user:pass@host.com:5432/path?k=v#f"
    F.Println(`[*]Source url:`, sUrl)

    u, err := url.Parse(sUrl)
    if err != nil {
        panic(err)
    }

    F.Println(`[*]Url.Scheme:`, u.Scheme)

    F.Println(`[*]Url.User:`, u.User)

    F.Println(`[*]Url.User.Username():`, u.User.Username())

    p, _ := u.User.Password()
    F.Println(`[*]Url.User.Password():`, p)
    F.Println(`[*]Url.Host:`, u.Host)

    host, port, _ := net.SplitHostPort(u.Host)
    F.Println(`[*]Url.Host.host:`, host)
    F.Println(`[*]Url.Host.Port:`, port)

    F.Println(`[*]Url.Path:`, u.Path)
    F.Println(`[*]Url.Fragment:`, u.Fragment)

    F.Println(`[*]Url.RawQuery:`, u.RawQuery)

    m, _ := url.ParseQuery(u.RawQuery)
    F.Println(`[*]Url.ParseQuery(Url.RawQuery):`, m)

    F.Println(`[*]Url.ParseQuery returns a type of url.Valuemap value: map[string][]string`)
}
