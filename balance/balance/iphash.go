package balance

import (
	"fmt"
	"math/rand"
	"hash/crc32"
	"time"
)

func init() {
	//RegisterBalancer(, &RandomBalance{})
	mBalanceManger["iphash"]=&HashBalance{}
}

type HashBalance struct {
	insts []*Instance
}
func (p *HashBalance)RegisterNodes(insts []*Instance)IBalance  {
	p.insts=insts
	return p
}

func (p *HashBalance)Add(node *Instance)error {
	p.insts=append(p.insts,node)
	return nil
}

func (p *HashBalance) DoBalance(key ...string) (inst *Instance, err error) {

	var defKey string
	if len(key) > 0 {
		defKey = key[0]
	}else{
		rand.Seed(time.Now().UnixNano())
		defKey= fmt.Sprintf("%d", rand.Int())
	}
	lens := len(p.insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	fmt.Println(hashVal,defKey)
	inst = p.insts[index]
	return
}
