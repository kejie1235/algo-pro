package main

import (
	"fmt"
	"crypto/md5"
	"hash/crc64"
	"common"
	"strconv"
)

func main()  {
	fmt.Printf("hello world!\n")

	testConsistentHash()
}

func testMd5Func() {
	testString := "func1235"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(testString))
	result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%d\n\n", result)
}

func testCrc32Func() {
	str := "123456asdfgg"

	result := crc64.Checksum([]byte(str), crc64.MakeTable(crc64.ISO))
	fmt.Printf("%d\n", result)
}

func testConsistentHash()  {
	
	impl := common.NewConsistentHashImp(20, func(p, q *common.HashCircleItem) bool {
		return p.Key < q.Key
	})

	serverList := []common.ServerNode{
		{"192.168.0.1", "Svr01"},
		{"192.168.0.2", "Svr02"},
		{"192.168.0.3","Svr03"},
		{"192.168.0.4", "Svr04"},
		{"192.168.0.5", "Svr05"},
		{"192.168.0.6", "Svr06"},
	}

	impl.Add(serverList)

	server := "127.0.0."

	statics := make(map[string]int)
	for index:=1; index<100; index++ {

		str := server + strconv.Itoa(index)
		node := impl.Get(str)

		if node != nil {
			fmt.Printf("[%s]节点的hash值为 %d, 被路由到的节点为 [%s]\n",
							str, crc64.Checksum([]byte(str), crc64.MakeTable(crc64.ECMA)), node.ServerIP)

			statics[node.ServerIP]++
		}
	}
	fmt.Printf("一致性Hash命中情况: %v", statics)
}
