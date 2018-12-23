package ttimer

import (
	"time"
	"fmt"
)

type Message struct {
	Type 		string
	Body		string
}
type HeapTimer struct {
	Expire  time.Time
	Timeout int64
	cb      func(v interface{})
	msg     *Message
}

func Callback(v interface{})  {
	msg := v.(Message)
	fmt.Printf("Msg Type: %s\n", msg.Type)
	fmt.Printf("Msg Type: %s\n", msg.Body)
}

func NewHeapTimer(timeout int64, callback func(v interface{}), msg_ *Message) *HeapTimer {
	this := new(HeapTimer)
	this.Expire = time.Now().Add(time.Duration(timeout))
	this.cb = callback
	this.Timeout = timeout
	this.msg = msg_
	return this
}
