package main

import (
	"fmt"
	"net/http"
	s "strings"
)

func main() {
	var r *http.Request

	fmt.Printf("Content-type: text/html\n\n")
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<title>login</title>")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	if s.Compare(password, username) == 0 {
		fmt.Println("<p>invalid username/password")
	}
}