package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

func josephusWithCList(l *common.List, m int) []interface{} {

	rel := make([]interface{}, 0, l.Size()-1)
	if l.Size() <= 1 {
		return []interface{}{l.Head().GetValue()}
	}

	cur := l.Head()

	var pre *common.ListElm

	count := 0

	for l.Size() != 1 {
		count ++
		if count == m {
			rel = append(rel, cur.GetValue())
			l.Rem_next(pre)
			count = 0
		}
		pre = cur
		cur = cur.Next()
	}
	return rel
}

func josephusNodes(l *common.ListElm, m int) []interface{} {
	//todo
	return nil
}

func main() {

	l1 := &common.List{}
	(*common.CList)(l1).Ins_next((*common.List)(l1).Tail(), 1)
	(*common.CList)(l1).Ins_next((*common.List)(l1).Tail(), 2)
	(*common.CList)(l1).Ins_next((*common.List)(l1).Tail(), 3)
	(*common.CList)(l1).Ins_next((*common.List)(l1).Tail(), 4)
	(*common.CList)(l1).Ins_next((*common.List)(l1).Tail(), 5)

	//l1.Traverse(common.Tf)

	rel := josephusWithCList(l1, 2)

	fmt.Println("rel ", rel) //2415

}
