package balance

import "errors"

func init()  {
	//RegisterBalancer("roundrobin", &RoundRobinBalance{})
	mBalanceManger["roundrobin"]=&RoundRobinBalance{}
}
type RoundRobinBalance struct {
	curIndex int
	insts []*Instance
}

func (p *RoundRobinBalance)RegisterNodes(insts []*Instance)IBalance  {
         p.insts=insts
	return p
}

func (p *RoundRobinBalance)Add(node *Instance)error{
	p.insts=append(p.insts,node)
	return nil
}

func (p *RoundRobinBalance)DoBalance(key ...string) (inst *Instance, err error) {
	if len(p.insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(p.insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = p.insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}