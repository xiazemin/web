package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const indexHtml = `<!DOCTYPE html>
<html>
<head><title>Go JSONP Server</title></head>
<body>
<button id="btn">Click to get HTTP header via JSONP</button>
<pre id="result"></pre>
<script>
'use strict';

var btn = document.getElementById("btn");
var result = document.getElementById("result");

function myCallback(acptlang) {
  result.innerHTML = JSON.stringify(acptlang, null, 2);
}

function jsonp() {
  result.innerHTML = "Loading ...";
  var tag = document.createElement("script");
  tag.src = "/jsonp?callback=myCallback";
  document.querySelector("head").appendChild(tag);
}

btn.addEventListener("click", jsonp);
</script>
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	callbackName := r.URL.Query().Get("callback")
	if callbackName == "" {
		fmt.Fprintf(w, "Please give callback name in query string")
		return
	}

	b, err := json.Marshal(r.Header)
	if err != nil {
		fmt.Fprintf(w, "json encode error")
		return
	}

	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, "%s(%s);", callbackName, b)
}

func main() {
	http.HandleFunc("/jsonp", jsonpHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8002", nil)
}
