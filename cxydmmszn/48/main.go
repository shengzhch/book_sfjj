package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

//链表回文结构判断

var cf common.CF = common.Compare_int

//用栈判断
func IsPalindrom1(list *common.List) bool {
	stack := common.Stack{}
	stack.Init()

	//把元素放入到栈中
	list.Traverse(func(e *common.ListElm, args ...interface{}) bool {
		stack.Push(e.GetValue())
		return false
	})
	var rel = new(bool)
	*rel = true

	//判断是否为回文，当不是时，停止遍历，写入false返回
	list.Traverse(func(e *common.ListElm, args ...interface{}) bool {
		if cm := cf(e.GetValue(), stack.PopValue()); cm != 0 {
			*args[0].(*bool) = false
			return true
		}
		return false
	}, rel)

	return *rel
}

//用栈判断
func IsPalindrom2(list *common.List) bool {
	stack := common.Stack{}
	stack.Init()

	//优化方向：无需完全压入栈，压入右半部分即可
	right := list.Head().Next()
	cur := list.Head()
	for cur.Next() != nil && cur.Next().Next() != nil {
		right = right.Next()    //移动一个
		cur = cur.Next().Next() //移动两个单位
	}

	for right != nil {
		stack.Push(right.GetValue())
		right = right.Next()
	}

	var rel = new(bool)
	*rel = true

	//判断是否为回文，当不是时，停止遍历，写入false返回
	list.Traverse(func(e *common.ListElm, args ...interface{}) bool {
		if stack.IsEmpty() {
			return true
		}
		if cm := cf(e.GetValue(), stack.PopValue()); cm != 0 {
			*args[0].(*bool) = false
			return true
		}
		return false
	}, rel)

	return *rel
}



func main() {
	l1 := &common.List{}
	l1.Init()
	l1.Ins_next(l1.Tail(), 1)
	l1.Ins_next(l1.Tail(), 2)
	l1.Ins_next(l1.Tail(), 3)
	l1.Ins_next(l1.Tail(), 2)
	l1.Ins_next(l1.Tail(), 1)

	fmt.Println("rel ", IsPalindrom2(l1))
}
