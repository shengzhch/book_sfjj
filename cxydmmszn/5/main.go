package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

type TwostackQueue struct {
	spush *common.Stack
	spop  *common.Stack
}

func (s *TwostackQueue) Init() {
	s.spush.Init()
	s.spop.Init()
}

func (s *TwostackQueue) Enqueue(data interface{}) {
	s.spush.Push(data)
}

func (s *TwostackQueue) Dequeue() interface{} {
	if s.spush.Size() == 0 && s.spop.Size() == 0 {
		return nil
	} else if s.spop.Size() == 0 { //注意只有当前的spop为空时，进行一次转栈操作
		for s.spush.Size() > 0 {
			s.spop.Push(s.spush.PopValue())
		}
	}

	return s.spop.PopValue()
}

func (s *TwostackQueue) GetHead() interface{} {
	//转栈操作
	if s.spop.Size() == 0 {
		for s.spush.Size() > 0 {
			s.spop.Push(s.spush.PopValue())
		}
	}

	return s.spop.Top()
}

func (s *TwostackQueue) Show() {
}

func main() {

	ms := &TwostackQueue{
		&common.Stack{}, &common.Stack{},
	}
	ms.Init()
	ms.Enqueue(1)
	ms.Enqueue(2)
	ms.Enqueue(3)
	ms.Enqueue(4)
	ms.Enqueue(5)
	v := ms.Dequeue()
	fmt.Println("v -- ", v)
	ms.Enqueue(6)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	ms.Enqueue(7)
	ms.Enqueue(8)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()

	ms.Enqueue(9)
	ms.Enqueue(10)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
	v = ms.Dequeue()
	fmt.Println("v -- ", v)
}
