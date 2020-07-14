package balance

import (

	"time"
	"errors"
	"math/rand"
)

func init() {
	//RegisterBalancer(, &RandomBalance{})
	mBalanceManger["random"]=&RandomBalance{}
}

type RandomBalance struct {
	insts []*Instance
}

func (p *RandomBalance)RegisterNodes(insts []*Instance)IBalance  {
	p.insts=insts
	return p
}

func (p *RandomBalance)Add(node *Instance)error{
	p.insts=append(p.insts,node)
	return nil
}

func (p *RandomBalance) DoBalance(key ...string) (inst *Instance, err error) {
	if len(p.insts)==0{
		err=errors.New("no instance")
		return
	}
	rand.Seed(time.Now().UnixNano())
	return p.insts[rand.Intn(len(p.insts))],nil
}
