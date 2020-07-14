package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"fmt"
)

type FastCGI struct{}

func (s *FastCGI) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c:="<html><head>Hello, fastcgi</head><body>Hello, fastcgi</body></html>"
	resp.Header().Set("Content-length",fmt.Sprintf("%d",len(c)))
	resp.Write([]byte(c))
	fmt.Println("cgi get input %#v",req)
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9000")
	srv := new(FastCGI)
	fcgi.Serve(listener, srv)
	select {}
}
