package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		time.Sleep(2*time.Second)
		fmt.Fprintln(w,"hello world")
	})
	http.ListenAndServe(":8889",nil)
}
