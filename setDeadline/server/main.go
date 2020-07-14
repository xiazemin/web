package main

import (
	"net/http"
	"time"
	"log"
)

func main() {
	srv:=&http.Server{
		Addr:":8876",
		ReadTimeout:1*time.Second,
		WriteTimeout:10*time.Second,
	}
	log.Println(srv.ListenAndServe())
}
