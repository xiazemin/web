package balance

import (
	"hash/crc32"
	"errors"
	//"github.com/golang/go/src/fmt"
)

const NumSLots  =16384

func init() {
	//RegisterBalancer(, &RandomBalance{})
	mBalanceManger["slotmaphash"]=&SlotMapHash{
		slots:make(map[*Instance]SlotRange),
	}

}

type SlotRange struct {
	min uint32
	max uint32
}
type SlotMapHash struct {
	insts []*Instance
	slots map[*Instance]SlotRange
}
func (p *SlotMapHash)RegisterNodes(insts []*Instance)IBalance  {
	p.insts=insts
	return p
}

func (p *SlotMapHash)Add(node *Instance)error {
	p.insts=append(p.insts,node)
	return nil
}

func (p *SlotMapHash)SetSlot(node *Instance,min,max uint32)error {
	//if _,ok:=p.slots[node];!ok{
	//	return  errors.New("no host added")
	//}
	p.slots[node]=SlotRange{
		min:min,
		max:max,
	}
     return nil
}

func (p *SlotMapHash) DoBalance(key ...string) (inst *Instance, err error) {
        ks:=""
	for _,k:=range key{
		ks+=k
	}
	slot:=p.crc32Hash(ks)%NumSLots
	//fmt.Println(p.slots,slot)
	for s,i:=range p.slots{
		//fmt.Println(i,slot)
		if slot>=i.min && slot<=i.max{
			return s,nil
		}
	}
	return nil,errors.New("slot not found")
}

func  (p *SlotMapHash) crc32Hash(key string)uint32{
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum([]byte(key), crc32q)
}

