package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

func findk(l, a, b int) int {
	r := float64(a) / float64(b)
	if r == 0 {
		return 0
	}


	fmt.Println("r ",r)

	for i := 1; i <= l; i++ {
		if r > float64(i)/float64(l) {
			continue
		}

		return i
	}

	return 0
}

//删除第K个节点
func deleteListKNode(l *common.List, k int) {
	if k < 1 || k > l.Size() {
		return
	}

	tar := k-1

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

func main() {
	l1 := &common.List{}
	l1.Init()
	l1.Ins_next(l1.Tail(), 1)
	l1.Ins_next(l1.Tail(), 2)
	l1.Ins_next(l1.Tail(), 3)
	l1.Ins_next(l1.Tail(), 4)
	l1.Ins_next(l1.Tail(), 5)

	deleteListKNode(l1, findk(l1.Size(), 11, 10))

	l1.Traverse(common.Tf)

}
