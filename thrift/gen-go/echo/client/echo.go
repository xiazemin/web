package main
import (
"fmt"
"log"
"net"
"os"

"git.apache.org/thrift.git/lib/go/thrift"
"github.com/xiazemin/thrift/gen-go/echo"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8000"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport:= transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:8000", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	req := &echo.EchoReq{
		Msg:"You are welcome.",
		Trace:"test",
		Help:&echo.Help{Time:1},
	}
	res, err := client.Echo(req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res.Msg)
	fmt.Println("well done")

}