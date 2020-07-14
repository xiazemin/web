package main

import "fmt"
import "os"

func main() {
	ip := os.Getenv("REMOTE_ADDR")
	fmt.Printf("Content-type: text/plain\n\n")
	fmt.Println(ip)
}
