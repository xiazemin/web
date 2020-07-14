package balance
type IBalance interface {
	RegisterNodes(insts []*Instance)IBalance
	//*IBalance is pointer to interface, not interface
	DoBalance(key ...string) (inst *Instance, err error)
	Add(node *Instance)error
}

var mBalanceManger map[string]IBalance

func init()  {
	mBalanceManger=make(map[string]IBalance)
}

func Get(name string)IBalance  {
	return mBalanceManger[name]
}
