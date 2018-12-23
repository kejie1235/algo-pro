package ttimer

import (
	"errors"
)

type HeapTimerManager struct {
	timers		[]*HeapTimer
	capacity	int
	cur			int
}

func NewHeapTimerManger(cap int) *HeapTimerManager {
	this := new(HeapTimerManager)
	this.capacity = cap
	this.cur = 0
	this.timers = make([]*HeapTimer, this.capacity)
	return this
}

func (this *HeapTimerManager) AddTimer(timeout int64, callback func(v interface{}), msg_ *Message)  {
	heapTimer := NewHeapTimer(timeout, callback, msg_)
	this.timers = append(this.timers, heapTimer)
}

func (this *HeapTimerManager)Empty() bool{
	return this.cur == 0
}

func (this *HeapTimerManager)TopTimer() (*HeapTimer,error) {
	if this.Empty() {
		return nil, errors.New("empty")
	}

	return this.timers[0], nil
}

func (this *HeapTimerManager)DelTimer(index int)  {
	this.timers = append(this.timers[:index], this.timers[index+1:]...)

}









