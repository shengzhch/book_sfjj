package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

func getCommonListData(l1, l2 *common.List, cf common.CF) []interface{} {
	s1 := l1.Head()
	s2 := l2.Head()
	rel := make([]interface{}, 0)
	for s1 != nil && s2 != nil {
		cm := cf(s1.GetValue(), s2.GetValue())
		if cm < 0 {
			s1 = s1.Next()
		} else if cm > 0 {
			s2 = s2.Next()
		} else {
			rel = append(rel, s1.GetValue())
			s1 = s1.Next()
			s2 = s2.Next()

		}
	}

	return rel
}

func main() {
	l1 := &common.List{}
	l1.Init()
	l1.Ins_next(l1.Tail(), 1)
	l1.Ins_next(l1.Tail(), 3)
	l1.Ins_next(l1.Tail(), 4)
	l1.Ins_next(l1.Tail(), 6)
	l1.Ins_next(l1.Tail(), 7)


	l2 := &common.List{}
	l2.Init()
	l2.Ins_next(l2.Tail(), 1)
	l2.Ins_next(l2.Tail(), 2)
	l2.Ins_next(l2.Tail(), 3)
	l2.Ins_next(l2.Tail(), 5)
	l2.Ins_next(l2.Tail(), 7)


	rel := getCommonListData(l1, l2, common.Compare_int)

	fmt.Println("rel ", rel)

}
