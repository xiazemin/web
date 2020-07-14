package main

import (
	"fmt"
	"os"
	"net/http/cgi"
)

func main() {
	httpReq, err := cgi.Request()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	r := httpReq.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Content-type: text/html\n\n")
	fmt.Printf("<!DOCTYPE html>\n")
	fmt.Printf("<p>username: %s\n", username)
	fmt.Printf("<p>password: %s\n", password)
}