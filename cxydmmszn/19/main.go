package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
)

func getMaxArrfromArr(arr []interface{}, w int, cf common.CF) []interface{} {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}

	//双向链表（这里使用的双向链表，其实单链表（包含头结点和尾节点）即可）
	//存储 数组的下标
	qmax := new(common.DList)
	qmax.Init()

	res := make([]interface{}, len(arr)-w+1)

	index := 0

	for i := 0; i < len(arr); i++ {

		//把小于等于当前值的元素全部移出链表
		for qmax.Size() > 0 && cf(arr[qmax.GetTailValue().(int)], arr[i]) <= 0 {
			qmax.Rem_ele(qmax.GetTail())
		}

		//插入当前值的下标
		qmax.Ins_next(qmax.GetTail(), i)

		//第i个元素位于框中，框中最左端元素的下线i-w+1,之前的元素需要移除
		if cf(qmax.GetHeadValue(), i-w) <= 0 {
			qmax.Rem_ele(qmax.GetHead())
		}

		//当框中元素完整是，取出此时的最大值，放入结果集中。整个赋值区间 w-1 ~ len(arr)，长度与res长度相等
		if (i >= w-1) {
			res[index] = arr[qmax.GetHeadValue().(int)]
			index++
		}
	}

	return res

}

func main() {
	arr := []interface{}{4, 3, 5, 4, 3, 3, 6, 7}
	rel := getMaxArrfromArr(arr, 3, common.Compare_int)
	fmt.Println("rel ", rel)
}
