package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/xiazemin/web/balance/balance"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for i:=0;i<8;i++{
		balanceName="slotmaphash"
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		fmt.Println(balance.Get(balanceName).Add(one))
		if smh,ok:=balance.Get(balanceName).(*balance.SlotMapHash);ok{
			fmt.Println(smh.SetSlot(one,uint32(balance.NumSLots/8*i),uint32(balance.NumSLots/8*(i+1))))
		}
	}
	for i:=0;i<8 ;i++  {
		balanceName="slotmaphash"
		fmt.Println(balance.Get(balanceName).DoBalance("192.168.1.2"))
		fmt.Println(balance.Get(balanceName).DoBalance("192.168.89.255"))
	}

	for i:=0;i<3;i++{
		balanceName="consitanthash"
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		fmt.Println(balance.Get(balanceName).Add(one))
		fmt.Println(balance.Get(balanceName).DoBalance("192.168.1.2"))
	}
	for i:=0;i<2;i++{
		balanceName="iphash"
		fmt.Println(balance.Get(balanceName).RegisterNodes(insts).DoBalance("192.168.1.2"))
		fmt.Println(balance.Get(balanceName).DoBalance("192.168.2.2"))

	}
	for i:=0;i<20;i++{
		balanceName="roundrobin"
		fmt.Println(balance.Get(balanceName).RegisterNodes(insts).DoBalance(balanceName))

	}
	balanceName = "random"
	for {
		inst, err := balance.Get(balanceName).RegisterNodes(insts).DoBalance(balanceName)
		if err != nil {
			fmt.Println("do balance err:", err)
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}