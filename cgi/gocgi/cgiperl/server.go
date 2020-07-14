package main

import(
	"net/http/cgi"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		handler := new(cgi.Handler)
		handler.Path = "/Users/didi/goLang/src/github.com/xiazemin/cgi/gocgi/cgiperl" + r.URL.Path
		log.Println(handler.Path)
		handler.Dir = "/Users/didi/goLang/src/github.com/xiazemin/cgi/gocgi/cgiperl"

		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8989",nil))
}