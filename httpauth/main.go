package main
import (
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	finalHandler := http.HandlerFunc(final)
	authHandler := httpauth.SimpleBasicAuth("username", "password")

	http.Handle("/httpauth", authHandler(finalHandler))
	http.ListenAndServe(":8089", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
