package main

import (
        "net/http"
        "io/ioutil"
        "fmt"
)

func main() {

        http.HandleFunc("/", serveStaticFile("public/index.html"))
        http.HandleFunc("/callback", serveStaticFile("public/callback.html"))
        http.HandleFunc("/js/app.js", serveStaticFile("public/resources/js/app.js"))

        listenErr := http.ListenAndServe(":8080", nil)

        if listenErr != nil {
                panic(fmt.Sprintf("Could not start server. Details: %s", listenErr.Error()))
        }
}

func serveStaticFile(htmlPath string) func(w http.ResponseWriter, r *http.Request) {

        return func(w http.ResponseWriter, r *http.Request) {
                w.Write(readFile(htmlPath))
        }
}

func readFile(filePath string) []byte {
        body, err := ioutil.ReadFile(filePath)
        if err != nil {
                panic("Could not read file " + filePath)
        }

        return body
}