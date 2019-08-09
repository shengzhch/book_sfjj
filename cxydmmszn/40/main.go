package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

func revarseList(l *common.List) {
	head := l.Head()
	tail := head
	var pre *common.ListElm

	for head != nil {
		next := head.Next()
		head.SetNext(pre)

		pre = head
		head = next
	}
	l.SetHead(pre)
	l.SetTail(tail)
}

func revarseDList(l *common.DList) {
	head := l.GetHead()
	tail := head
	var pre *common.DListElm

	for head != nil {
		next := head.Next()
		head.SetNext(pre)
		head.SetPre(next)

		pre = head
		head = next
	}
	l.SetHead(pre)
	l.SetTail(tail)
}

func main() {

	l1 := &common.DList{}
	l1.Init()
	l1.Ins_next(l1.GetTail(), 1)
	l1.Ins_next(l1.GetTail(), 2)
	l1.Ins_next(l1.GetTail(), 3)
	l1.Ins_next(l1.GetTail(), 4)
	l1.Ins_next(l1.GetTail(), 5)

	revarseDList(l1)
	fmt.Println("head -",l1.GetHead().GetValue())
	fmt.Println("tail -",l1.GetTail().GetValue())

	l1.Traverse(common.Tdf)

	l1.ReTraverse(common.Tdf)
}
