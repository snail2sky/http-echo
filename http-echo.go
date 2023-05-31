package main

import (
        "flag"
        "fmt"
        "log"
        "net/http"
)

var listenHost = flag.String("host", "0.0.0.0", "HTTP echo server listen on this address.")
var listenPort = flag.Int("port", 8080, "HTTP echo server listen on this port.")

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {
        log.Println(fmt.Sprintf("HTTP echo server handle request: <%s %s %s> from client (%s) ", request.Method, request.RequestURI, request.Proto, request.RemoteAddr))
        err := request.Write(writer)
        if err != nil {
                log.Println(err)
        }
}

func main() {
        flag.Parse()
        var listenOn = fmt.Sprintf("%s:%d", *listenHost, *listenPort)
        log.Println("HTTP echo server listen on:", listenOn)

        http.HandleFunc("/", EchoHandler)
        log.Println(http.ListenAndServe(listenOn, nil))
}