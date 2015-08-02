package main

import (
    "fmt"
    "log"
    "io"
    "os"
    "regexp"
    "net/url"
    "net/http"
)

var (
    Access   *log.Logger
)

func InitLogger() {
    var AccessFile = "access.log"
    file, err := os.OpenFile(AccessFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", AccessFile, ":", err)
    }

    multi := io.MultiWriter(file, os.Stdout)

    Access = log.New(multi, "", log.Ldate|log.Ltime)
}

var UrlPattern = regexp.MustCompile(`^https?://.*$`)

var logFilePath = "./access.log"

func handler(w http.ResponseWriter, r *http.Request) {
    params, _ := url.ParseQuery(r.URL.RawQuery)
    
    var urlParam = params.Get("url")
    Access.Println(r.RemoteAddr, "accessing url:", urlParam)

    if UrlPattern.MatchString(urlParam) {
        http.Redirect(w, r, urlParam, 302)
    } else {
        fmt.Fprintf(w, "The url is invalid")
    }
}

func main() {
    InitLogger()
    http.HandleFunc("/r", handler)
    http.ListenAndServe(":8080", nil)
}
