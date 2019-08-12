package main

import (
	"github.com/shengzhch/book_sfjj/common"
)

//部分反转

func Reversepart(l *common.List, from, to int) {
	len := 0

	//范围的前一个节点和后一个节点
	var fpre, tPos *common.ListElm
	var node1 *common.ListElm

	node1 = l.Head()

	for node1 != nil {
		len++
		if len == from-1 {
			fpre = node1
		}

		if len == to+1 {
			tPos = node1
		}
		node1 = node1.Next()
	}

	if fpre == nil {
		node1 = l.Head()
	} else {
		node1 = fpre.Next()
	}

	node2 := node1.Next()

	node1.SetNext(tPos)

	//注意更新链表的尾部
	if tPos == nil {
		l.SetTail(node1)
	}

	var next *common.ListElm
	//范围内反转
	for node2 != tPos {
		next = node2.Next()
		node2.SetNext(node1)
		node1 = node2
		node2 = next
	}

	if fpre != nil {
		fpre.SetNext(node1)
		return
	}
	l.SetHead(node1)
}

func main() {
	l1 := &common.List{}
	l1.Init()
	l1.Ins_next(l1.Tail(), 1)
	l1.Ins_next(l1.Tail(), 2)
	l1.Ins_next(l1.Tail(), 3)
	l1.Ins_next(l1.Tail(), 4)
	l1.Ins_next(l1.Tail(), 5)

	Reversepart(l1, 1, 5)
	l1.Traverse(common.Tf)
}
