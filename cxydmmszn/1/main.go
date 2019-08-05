package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

type Mystack struct {
	data *common.Stack
	min  *common.Stack
}

func (s *Mystack) Init() {
	s.data.Init()
	s.min.Init()
}

func (s *Mystack) Push(data interface{}, cf common.CF) {
	if s.min.Size() == 0 {
		s.min.Push(data)
	} else if cf(data, s.Getmin()) <= 0 {
		s.min.Push(data)
	}

	s.data.Push(data)
}

func (s *Mystack) Pop(cf common.CF) interface{} {
	pv := s.data.PopValue()
	if cf(pv, s.Getmin()) == 0 {
		s.min.PopValue()
	}
	return pv
}

func (s *Mystack) Getmin() interface{} {
	return s.min.Top()
}

func (s *Mystack) Show() {
	fmt.Println("data --- ")

	tf := func(e *common.ListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}

	(*common.List)(s.data).Traverse(tf)

	fmt.Println("min --- ")
	(*common.List)(s.min).Traverse(tf)
}

func main() {

	cf := common.Compare_int
	ms := &Mystack{
		&common.Stack{}, &common.Stack{},
	}
	ms.Init()

	ms.Push(1, cf)
	ms.Push(2, cf)
	ms.Push(1, cf)
	ms.Push(5, cf)
	ms.Push(4, cf)
	ms.Push(3, cf)

	ms.Show()
	ms.Pop(cf)
	ms.Pop(cf)
	ms.Show()
	ms.Pop(cf)
	ms.Pop(cf)
	ms.Show()
}
