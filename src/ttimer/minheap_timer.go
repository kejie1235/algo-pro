package ttimer

import (
	"time"
	"errors"
)

type HeapTimerManager struct {
	timers		[]HeapTimer
	capacity	int
	cur			int
}

func NewHeapTimer(timeout int64, callback func(v interface{}), msg_ *Message)  {
	this := new(HeapTimer)
	this.Now = time.Now()
	this.cb = callback
	this.Timeout = timeout
	this.msg = msg_
}

func NewHeapTimerManger(cap int) {
	this := new(HeapTimerManager)
	this.capacity = cap
	this.cur = 0
	this.timers = make([]HeapTimer, this.capacity)
}

func (this *HeapTimerManager) AddTimer()  {
	
}

func (this *HeapTimerManager)Empty() bool{
	return this.cur == 0
}

func (this *HeapTimerManager)TopTimer() (*HeapTimer,error) {
	if this.Empty() {
		return nil, errors.New("empty")
	}

	return &this.timers[0], nil
}










