package main

import (
	"fmt"
	"net"
	"strconv"
	"os"
	"time"
	"strings"
	"github.com/xiazemin/gofcgid/fcgiclient"
)

func main() {
	service := "0.0.0.0:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	fmt.Println("Listen on: ", tcpAddr.IP.String(), ":", tcpAddr.Port)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Accept from: ", conn.RemoteAddr().String())
		
		go handleClient(conn)
		fmt.Println("after close connect")
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // set 2 minutes timeout
	request := make([]byte, 2048)                         // set maxium request length to 128KB to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		fmt.Println("continue read")
		read_len, err := conn.Read(request)
		fmt.Println("continue read finish")
		if err != nil {
			fmt.Println(err,"502 bad gateway")
			conn.Write([]byte("\r\n"))
			break
		}
		fmt.Println(read_len)
		if read_len == 0 {
			fmt.Println(read_len)
			break // connection already closed by client
		}  else {
			content := strings.Trim(string(request[:read_len]), "\r\n")
			fmt.Println("We get ", read_len, " bytes: ", content)
			conn.Write(fcgiProcessFunc(content))
		}

		request = make([]byte, 2048) // clear last read content
	}
	fmt.Println("before close connect")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func fcgiProcessFunc(value string) []byte {

	postParams  := value
	queryParams := "foo=bar&hello=world"

	env := make(map[string]string)
	env["REQUEST_METHOD"] = "POST"
	env["SCRIPT_FILENAME"] = "/Users/guweigang/wwwroot/test.php"
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["SERVER_PROTOCOL"] = "HTTP/1.1"
	env["QUERY_STRING"] = queryParams
	env["CONTENT_TYPE"] = "application/x-www-form-urlencoded"
	env["CONTENT_LENGTH"] = strconv.FormatInt(int64(len(postParams)), 10)

	fcgi, err := fcgiclient.New("127.0.0.1", 9000)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	content, err := fcgi.Request(env, postParams)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println("before content\r\n")
	fmt.Println(string(content))
	fmt.Println("after content\r\n")
	return content;

	// data := strings.SplitAfter(string(content[:]), "\r\n\r\n")
	// fmt.Printf("%v\n", data[1])
	// fmt.Printf("content: %s", content)
}
