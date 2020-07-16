package main

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"
)

func goppf() {
	fm, err := os.OpenFile("./tmp/mem.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(fm)
	fm.Close()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}
