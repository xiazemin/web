// file name: server.go

package main

import (
	"os"
	"io"
	"net"
	"log"
	"strconv"
	// "time"
)

// 把接收到的内容append到文件
func writeFile(content []byte) {
	if len(content) != 0 {
		fp, err := os.OpenFile("test_1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		defer fp.Close()
		if err != nil {
			log.Fatalf("open file faild: %s\n", err)
		}
		_, err = fp.Write(content)
		if err != nil {
			log.Fatalf("append content to file faild: %s\n", err)
		}
		log.Printf("append content: 【%s】 success\n", string(content))
	}
}

// 获取已接收内容的大小
// (断点续传需要把已接收内容大下通知客户端从哪里开始发送文件内容)
func getFileStat() int64 {
	fileinfo, err := os.Stat("test_1.txt")
	if err != nil {
		// 如果首次没有创建test_1.txt文件，则直接返回0
		// 告诉客户端从头开始发送文件内容
		if os.IsNotExist(err) {
			log.Printf("file size: %d\n", 0)
			return int64(0)
		}
		log.Fatalf("get file stat faild: %s\n", err)
	}
	log.Printf("file size: %d\n", fileinfo.Size())
	return fileinfo.Size()
}

func serverConn(conn net.Conn) {
	defer conn.Close()
	for {
		var buf = make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("server io EOF\n")
				return
			}
			log.Fatalf("server read faild: %s\n", err)
		}
		log.Printf("recevice %d bytes, content is 【%s】\n", n, string(buf[:n]))
		// 判断客户端发送过来的消息
		// 如果是'start-->‘则表示需要告诉客户端从哪里开始读取文件数据发送
		switch string(buf[:n]) {
		case "start-->":
			off := getFileStat()
			// int conver string
			stringoff := strconv.FormatInt(off, 10)
			_, err = conn.Write([]byte(stringoff))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}
			continue
		case "<--end":
			// 如果接收到客户端通知所有文件内容发送完毕消息则退出
			log.Fatalf("receive over\n")
			return
		// default:
		//  time.Sleep(time.Second * 1)
		}
		// 把客户端发送的内容保存到文件
		writeFile(buf[:n])
	}
}

func main() {
	// 建立监听
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("error listen: %s\n", err)
	}
	defer l.Close()

	log.Println("waiting accept.")
	// 允许客户端连接，在没有客户端连接时，会一直阻塞
	conn, err := l.Accept()
	if err != nil {
		log.Fatalf("accept faild: %s\n", err)
	}
	serverConn(conn)
}