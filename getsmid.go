package main

import (
    "fmt"
    "encoding/xml"
    "net/http"
    "io/ioutil"
    "log"
    "time"
    "os"
    "strconv"
)

type Nicovideo struct {
    Link []string `xml:"channel>item>link"`
}

func main() {
    tag := os.Args[1]
    start, _ := strconv.Atoi(os.Args[2])
    end, _ := strconv.Atoi(os.Args[3])

    for i := start; i < end+1; i++ {
        url := "http://www.nicovideo.jp/tag/" + tag + "?sort=f&rss=2.0&page=" + strconv.Itoa(i)
        nc, err := getNico(url)

        if err != nil {
            log.Fatalf("Log: %v", err)
            return
        }

        for _, v := range nc.Link {
            fmt.Printf("%s\n", v)
        }
        time.Sleep(3 * time.Second)
    }
}

func getNico(feed string) (p *Nicovideo, err error) {
    res, err := http.Get(feed)
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