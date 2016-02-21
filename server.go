package main

import (
    //"fmt"
    "encoding/json"
    "io/ioutil"

    "net/http"
    "net/http/httputil"
)

func main() {
    director := func(req *http.Request) {
        url := "https://status.github.com/api/status.json"

        resp, _ := http.Get(url)
        defer resp.Body.Close()

        var status map[string]string

        content, _ := ioutil.ReadAll(resp.Body)

        _ = json.Unmarshal(content, &status)

        target := "github.com"
        req.URL.Scheme = "https"

        if status["status"] != "good" {
            target = "cache.gitstash.net"
            req.URL.Scheme = "http"
        }
        req.URL.Host = target
    }

    proxy := &httputil.ReverseProxy{Director: director}

    http.ListenAndServe(":8081", proxy)
}
