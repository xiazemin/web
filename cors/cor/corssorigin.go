package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
)

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func main() {
	http.HandleFunc("/", Entrance)
	http.HandleFunc("/ajax", TestCrossOrigin)
	/**
$ curl -i http://127.0.0.1:8000/ajax
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json; charset=utf-8
Date: Sun, 22 Sep 2019 13:51:51 GMT
Content-Length: 38

{"name":"benben_2015","msg":"success"}
	 */
	http.ListenAndServe(":8001", nil)
}

func Entrance(w http.ResponseWriter, r *http.Request) {
	t,_:=template.ParseFiles("/Users/didi/goLang/src/github.com/xiazemin/cors/templates/ajax.html")
	t.Execute(w, nil)
}

func TestCrossOrigin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var message Message
		message.Name = "benben_2015"
		message.Msg = "success"

		result, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		ResponseWithOrigin(w, r, http.StatusOK, result)
		return
	}
}
func ResponseWithOrigin(w http.ResponseWriter, r *http.Request, code int, json []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
	//"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	w.WriteHeader(code)
	w.Write(json)
}