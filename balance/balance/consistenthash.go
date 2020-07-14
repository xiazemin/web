package balance

import (
	"errors"
	"sort"
	"crypto/sha1"
	"hash/crc32"
	"fmt"
)

func init() {
	//RegisterBalancer(, &RandomBalance{})
	mBalanceManger["consitanthash"] = &ConsistantHashBalance{
		numOfVirtualNode: 20,
		circle:           make(map[uint32]*Instance),
		nodes:            make(map[*Instance]bool),
	}
}

type ConsistantHashBalance struct {
	numOfVirtualNode int  //每个节点对应的虚拟节点数
	hashSortedNodes  []uint32  //所有节点(真实、虚拟)hash 后得到的环,(顺时针排序)
	circle           map[uint32]*Instance //环上位置到节点映射,虚拟节点存在所以可能多个位置映射到同一节点

	nodes            map[*Instance]bool   //节点是否服务中
	insts            []*Instance   //节点列表
}

func (p *ConsistantHashBalance) RegisterNodes(insts []*Instance) IBalance {
	p.insts = insts
	return p
}

//add the node
func (p *ConsistantHashBalance) Add(node *Instance) error {
	if _, ok := p.nodes[node]; ok {
		return errors.New("host already existed")
	}
	p.nodes[node] = true
	// add virtual node
	for i := 0; i < p.numOfVirtualNode; i++ {
		virtualKey := p.getVirtualPosition(i, node)
		p.circle[virtualKey] = node
		p.hashSortedNodes = append(p.hashSortedNodes, virtualKey)
	}

	sort.Slice(p.hashSortedNodes, func(i, j int) bool {
		return p.hashSortedNodes[i] < p.hashSortedNodes[j]
	})
	return nil
}

//remove the node
func (p *ConsistantHashBalance) Remove(node*Instance) error {
	if _, ok := p.nodes[node]; !ok {
		return errors.New("host not exist")
	}
	p.nodes[node] = false
	return nil
}

func (p *ConsistantHashBalance) DoBalance(key ...string) (inst *Instance, err error) {
	if len(p.nodes) == 0 {
		return nil, errors.New("no host added")
	}
	if len(key)==0{
		return nil,errors.New("request host nill")
	}
	pos:=p.getRequestPosition(key)
	nearbyIndex := p.searchNearbyIndex(pos)
	nearHost := p.circle[p.hashSortedNodes[nearbyIndex]]
	return nearHost, nil
}

//节点(虚拟、真实)和请求key 都经过同一套hash算法,映射到2^32-1的环形空间,返回值就是位置
func (p *ConsistantHashBalance)genPosition(key string) uint32 {
	hash := sha1.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	crcTable := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum([]byte(hashBytes), crcTable)
}

//可以使用服务器的IP地址或者主机名作为关键字,加上虚拟节点，并且是按照顺时针排列：
func (p *ConsistantHashBalance) getVirtualPosition(i int, node *Instance) uint32 {
	key:=fmt.Sprintf("%d%s",i,fmt.Sprint(node))
	return p.genPosition(key)
}

func (p *ConsistantHashBalance) getRequestPosition(key []string) uint32 {
	ks:=""
	for _,k:=range key{
		ks+=k
	}
	return p.genPosition(ks)
}



func (p *ConsistantHashBalance) searchNearbyIndex(key uint32)uint32{
	//hashSortedNodes 是升序排序,找到第一个比当前位置大的点即可,(顺时针排序,左侧第一个就是机器节点)
	// Search uses binary search to find and return the smallest index i
	// in [0, n) at which f(i) is true
	i := sort.Search(len(p.hashSortedNodes), func(i int) bool {
		//如果节点被删除了,走下一个
		node:=p.circle[p.hashSortedNodes[i]]
		return p.hashSortedNodes[i]<=key && p.nodes[node]})
	if i == len(p.hashSortedNodes) {
		i = 0
	}
	return uint32(i)
}