package common

import (
	"strconv"
	"hash/crc64"
	"sort"
	"fmt"
)

/*
	implement consistent_hash by golang

	Monotonicity means that if something has been hashed into the corresponding buffer, a new buffer is added to the system.
		The result of the hash should be such that the original allocated content can be mapped to the new buffer
			without being mapped to other buffers in the old buffer set.
 */

type ServerNode struct {
	ServerIP		string
	ServerName		string
}

type HashCircleItem struct{
	Key  uint64
	Node ServerNode
}

type HashCircleCmpWrapper struct {
	HashCircle []HashCircleItem
	Cmp func(p, q *HashCircleItem) bool
}

func (wrapper HashCircleCmpWrapper) Len() int {
	return len(wrapper.HashCircle)
}

func (wrapper HashCircleCmpWrapper) Swap(i ,j int) {
	wrapper.HashCircle[i], wrapper.HashCircle[j] = wrapper.HashCircle[j], wrapper.HashCircle[i]
}

func (wrapper HashCircleCmpWrapper) Less(i,j int) bool{
	return wrapper.Cmp(&wrapper.HashCircle[i], &wrapper.HashCircle[j])
}

type ConsistentHashImp struct {
	VirtualNodeNumber int
	Wrapper           HashCircleCmpWrapper
}

func NewConsistentHashImp(virtualNodeNumber int, cmp func(p, q *HashCircleItem) bool) *ConsistentHashImp{
	this := new(ConsistentHashImp)
	this.VirtualNodeNumber = virtualNodeNumber
	this.Wrapper = HashCircleCmpWrapper{}
	this.Wrapper.Cmp = cmp
	return this
}

func (this *ConsistentHashImp)sortHashCircleByCryptoKey() {
	sort.Sort(this.Wrapper)
}

func (this *ConsistentHashImp)Add(serverList []ServerNode) {
	for _, server := range serverList {
		for repIndex := 1; repIndex <= this.VirtualNodeNumber; repIndex++ {
			crypto := server.ServerIP + "#" + strconv.Itoa(repIndex)
			key := crc64.Checksum([]byte(crypto), crc64.MakeTable(crc64.ECMA))

			fmt.Printf("虚拟节点[%s]被添加, hash值为%d\n", crypto, key)
			item := HashCircleItem{
				key,
				server,
			}
			this.Wrapper.HashCircle = append(this.Wrapper.HashCircle, item)
		}
	}

	this.sortHashCircleByCryptoKey()
	//fmt.Print(sort.IsSorted(this.Wrapper))
	//fmt.Print(this.Wrapper.HashCircle)
}


func (this *ConsistentHashImp)Get(key string) (*ServerNode) {
	cryptoKey := crc64.Checksum([]byte(key), crc64.MakeTable(crc64.ECMA))

	for _, item := range this.Wrapper.HashCircle {
		if item.Key > cryptoKey {
			return &item.Node
		}
	}
	return nil
}

