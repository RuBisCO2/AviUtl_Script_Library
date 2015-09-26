package main

import (
    "fmt"
    "encoding/xml"
    "net/http"
    "io/ioutil"
    "log"
    "os"
    "io"
    "bufio"
    "time"
)

type Nicovideo struct {
    Title string `xml:"thumb>title"`
    Desc string `xml:"thumb>description"`
    Fr string `xml:"thumb>first_retrieve"`
    Nn string `xml:"thumb>user_nickname"`
}

func main() {
    fp, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
        os.Exit(1)
    }
    defer fp.Close()

    rd := bufio.NewReaderSize(fp, 4096)
    for {
        line, _, err := rd.ReadLine()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Fprintf(os.Stderr, "%v", err)
            os.Exit(1)
        }
        url := "http://ext.nicovideo.jp/api/getthumbinfo/" + string(line)
        nc, err := getNico(url)

        if err != nil {
            log.Fatalf("Log: %v", err)
            return
        }
        fmt.Printf("%s,%s,%s,%s\n", nc.Title, nc.Nn, nc.Desc, nc.Fr)
        time.Sleep(3 * time.Second)
    }
}

func getNico(api string) (p *Nicovideo, err error) {
    res, err := http.Get(api)
    if err != nil {
        return nil, err
    }

    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    nc := new(Nicovideo)
    err = xml.Unmarshal(b, &nc)

    return nc, err
}