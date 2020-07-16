package main

import (
	"fmt"
	"net/http"
)

func welComeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "输入关键词，查看如何排查")
	fmt.Println("输入关键词，查看如何排查")
}

func main() {
	http.HandleFunc("/index", welComeHandler)
	fmt.Println(http.ListenAndServe(":8089", nil))
}
