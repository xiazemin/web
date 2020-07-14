package main

import (
	"fmt"

	"github.com/xiazemin/thrift/gen-go/echo"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type EchoServerImp struct {
}

func (e *EchoServerImp) Echo(req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())

	res := &echo.EchoRes{
		Msg: req.GetMsg(),
	}

	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":8000")
	if err != nil {
		panic(err)
	}

	processor := echo.NewEchoProcessor(&EchoServerImp{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
