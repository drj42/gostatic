package main

import (
        "log"
        "net/http"
        "flag"
)

var root, port, host string

func init() {
        flag.StringVar(&root, "root", ".", "Specify the directory to serve files from.  Default is the current working directory")
        flag.StringVar(&port, "port", "8080", "Specify server port.  Default is 8080")
        flag.StringVar(&host, "host", "127.0.0.1", "Specify host or IP. Default is 127.0.0.1")
        flag.Parse()
}


func main() {
        http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
                log.Printf("%v %v", r.Method, r.URL.String())
                http.FileServer(http.Dir(root)).ServeHTTP(w,r)
                return
        })

        log.Printf("Initializing static web server on %v:%v in %v\n", host, port, root)
        log.Fatal(http.ListenAndServe(host + ":" + port , nil))
}
