package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/xiazemin/web/UDP/config"
)

func main() {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		// Here must use make and give the lenth of buffer
		data := make([]byte, config.SERVER_RECV_LEN)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		strData := string(data)
		fmt.Println("Received:", strData)

		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Send:", upper)
	}

}
/*
代码分为三大块，第一部分使用ResolveUDPAddr函数构造服务端的地址信息，第二部分使用ListenUDP函数，对config,go中指定的端口和地址进行监听，第三部分是一个for循环，该循环不断地从UDP读取客户端数据，然后将其转换为大写字母返回给客户端。ReadFromUDP函数读取UDP数据包，如果没有数据可读，该函数会阻塞。
 */