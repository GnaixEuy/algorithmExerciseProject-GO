package main

import "container/list"

type MyQueue struct {
	l    list.List
	size int
}

func Constructor232() MyQueue {
	k := MyQueue{size: 0}
	return k
}

func (this *MyQueue) Push(x int) {
	this.l.PushBack(x)
	this.size++
}

func (this *MyQueue) Pop() int {
	v := this.l.Front()
	if v != nil {
		this.l.Remove(v)
		ret := v.Value
		return ret.(int)
	}
	return 0
}

func (this *MyQueue) Peek() int {
	front := this.l.Front()
	if front != nil {
		return front.Value.(int)
	} else {
		return 0
	}
}

func (this *MyQueue) Empty() bool {
	front := this.l.Front()
	if front != nil {
		return false
	}
	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
