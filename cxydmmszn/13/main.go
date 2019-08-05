package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

var (
	tf = func(e *common.ListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}

	cfi = common.Compare_int
)

func SortStack(s *common.Stack, cf common.CF) {
	help := new(common.Stack)

	//一个元素只能在help中或者s中，help中的元素始终是有序的从大到小
	for s.Size() > 0 {
		top := s.PopValue()

		//help 一直出栈，知道找到小于当前值的元素，出栈的元素加入到s中
		for help.Size() > 0 && cf(help.Top(), top) > 0 {
			s.Push(help.PopValue())
		}

		//将栈顶元素押入栈中
		help.Push(top)
	}

	for help.Size() > 0 {
		s.Push(help.PopValue())
	}

}
func main() {
	s := new(common.Stack)
	s.Push(3)
	s.Push(4)
	s.Push(1)
	s.Push(2)
	s.Push(5)


	SortStack(s,cfi)
	(*common.List)(s).Traverse(tf)

	fmt.Println()

}
