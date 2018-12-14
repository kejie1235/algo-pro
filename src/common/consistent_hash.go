package common

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


type ConsistentHash struct {
	VirtualNodeNumber	int
	HashCircle 			map[int64]*ServerNode
	HashFunc			func(key string) int64
}

func NewConsistentHash(virtualNodeNumber int) {
	this := new(ConsistentHash)
	this.VirtualNodeNumber = virtualNodeNumber
	this.HashCircle = make(map[int64] *ServerNode)
	//this.HashFunc = hash.
}

func NewConsistentHashByHash(virtualNodeNumber int, ob func(key string) int64, ){
	this := new(ConsistentHash)
	this.VirtualNodeNumber = virtualNodeNumber
	this.HashFunc = ob
}

func Add(serverList []ServerNode) {

}


func Get(key string) {

}


