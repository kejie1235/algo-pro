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
	Now			time.Time
	Timeout 	int64
	cb			func(v interface{})
	msg			*Message
}

func Callback(v interface{})  {
	msg := v.(Message)
	fmt.Printf("Msg Type: %s\n", msg.Type)
	fmt.Printf("Msg Type: %s\n", msg.Body)
}
