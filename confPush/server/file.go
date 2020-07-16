package server

import (
	"fmt"
	"net/http"
)

func Download() {
	h := http.FileServer(http.Dir("/Users/didi/PhpstormProjects/go/"))
	err := http.ListenAndServe(":9090", h)
	fmt.Println(err)
}
