package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

//单链表删除倒数第K个节点
func deleteListKNode(l *common.List, k int) {
	if k > l.Size() {
		return
	}

	tar := l.Size() - k

	var tarn *common.ListElm
	for i := 0; i < tar; i++ {
		if i == 0 {
			tarn = l.Head()
		} else {
			tarn = tarn.Next()
		}
	}

	l.Rem_next(tarn)
}

func deleteDListKNode(l *common.DList, k int) {
	if k > l.Size() {
		return
	}

	tar := l.Size() - k

	var tarn *common.DListElm
	tarn = l.GetHead()
	for i := 0; i < tar; i++ {
		tarn = tarn.Next()
	}

	l.Rem_ele(tarn)
}

func main() {
	l1 := &common.List{}
	l1.Init()
	l1.Ins_next(l1.Tail(), 1)
	l1.Ins_next(l1.Tail(), 2)
	l1.Ins_next(l1.Tail(), 3)
	l1.Ins_next(l1.Tail(), 4)
	l1.Ins_next(l1.Tail(), 5)

	deleteListKNode(l1, 2)

	l1.Traverse(common.Tf)

	l2 := &common.DList{}
	l2.Init()
	l2.Ins_next(l2.GetTail(), 1)
	l2.Ins_next(l2.GetTail(), 2)
	l2.Ins_next(l2.GetTail(), 3)
	l2.Ins_next(l2.GetTail(), 4)
	l2.Ins_next(l2.GetTail(), 5)

	fmt.Println("正序")

	deleteDListKNode(l2, 1)

	l2.Traverse(common.Tdf)

	fmt.Println("倒序")

	l2.ReTraverse(common.Tdf)

}
