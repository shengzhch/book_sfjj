package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

var (
	tf = func(e *common.ListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}
)

// 递归调用取出插入,最终取出栈底元素
func getAndRemoveLastEle(s *common.Stack) interface{} {
	rel := s.PopValue()
	if s.Size() == 0 {
		return rel
	} else {
		last := getAndRemoveLastEle(s)
		s.Push(rel)
		return last
	}
}

//递归调用反转插入,最终实现stack的逆序
func ReverseStack(s *common.Stack) {
	if s.Size() == 0 {
		return
	} else {
		data := getAndRemoveLastEle(s)
		ReverseStack(s)
		s.Push(data)
	}
}

func main() {

	ms := &common.Stack{}
	ms.Init()

	ms.Push(1)
	ms.Push(2)
	ms.Push(3)
	ms.Push(4)
	ms.Push(5)

	fmt.Println("before ------- ")
	(*common.List)(ms).Traverse(tf)

	ReverseStack(ms)
	fmt.Println("after ------- ")

	(*common.List)(ms).Traverse(tf)
}
