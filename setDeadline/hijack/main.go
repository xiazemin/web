package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
)
func main() {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
			return
		}
		conn, bufrw, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Don't forget to close the connection:
		defer conn.Close()
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		conn.SetReadDeadline(time.Now().Add(1*time.Second))
		bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
		time.Sleep(10*time.Second)
		bufrw.Flush()
		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error reading string: %v", err)
			return
		}
		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
		bufrw.Flush()
	})
	http.ListenAndServe(":8877",nil)
}