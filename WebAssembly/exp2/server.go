package main

import (
	"flag"
	"log"
	"net/http"
	"go/src/fmt"
)

var (
	listen = flag.String("listen", ":8089", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	*dir="/Users/didi/goLang/src/github.com/xiazemin/WebAssembly/exp2"
	fmt.Println(http.Dir(*dir))
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}