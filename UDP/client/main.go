package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/xiazemin/web/UDP/config"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	serverAddr := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen - written > config.SERVER_RECV_LEN {
				toWrite = line[written : written + config.SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			checkError(err)

			fmt.Println("Write:", toWrite)

			msg := make([]byte, config.SERVER_RECV_LEN)
			n, err = conn.Read(msg)
			checkError(err)

			fmt.Println("Response:", string(msg))
		}
	}

}

/*
使用Dial函数可以很方便的建立客户端的连接，代码主要部分在for循环，不断地从标准输入读取数据，按照服务器一次接收的数据情况，分批发送数据，然后读取服务器的响应。
 */