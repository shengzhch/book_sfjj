package main

import (
	"github.com/shengzhch/book_sfjj/common"
	"fmt"
)

//原则
//每一个数的父节点都是左边第一个比他大和右边第一个比他大的数中较小的一个
//若是左右两边都没有比他大的数，则该数为头结点

//生成是是一个树 （容易证明，数值不重复，最大值唯一）

//树的节点最多有两个孩子 即为二叉树( 节点的一侧 最多只要一个子节点)
//反证法 假设a有两个节点 不妨设 k1,k2且a>k1>k2,按照规则，则a不可能是k2的父节点


//
func getMaxTreeFromArr(arr []interface{}, cf common.CF) *common.BiTreeNode {
	nodes := make([]*common.BiTreeNode, len(arr))

	for i := 0; i < len(arr); i++ {
		nodes[i] = common.NewBiTreeNode(arr[i])
	}

	lmap := make(map[*common.BiTreeNode]*common.BiTreeNode)
	rmap := make(map[*common.BiTreeNode]*common.BiTreeNode)


	s := new(common.Stack)


	//每个节点对应左边第一个大于它的节点
	for i := 0; i < len(nodes); i++ {
		curNode := nodes[i]


		//一直找到栈顶元素大于当前等于当前值
		for s.Size() > 0 && cf(s.Top().(*common.BiTreeNode).Data(), curNode.Data()) < 0 {
			popStackSetMap(s, lmap)
		}
		s.Push(curNode)
		//(*common.List)(s).Traverse(common.Tf)
		//fmt.Println("-----------")
	}
	//(*common.List)(s).Traverse(common.Tf)
	//fmt.Println("-----------")

	for !s.IsEmpty() {
		popStackSetMap(s, lmap)
	}


	//每个节点对应左边第一个大于它的节点
	for i := len(nodes) - 1; i >= 0; i-- {
		curNode := nodes[i]

		for s.Size() > 0 && cf(s.Top().(*common.BiTreeNode).Data(), curNode.Data()) < 0 {
			popStackSetMap(s, rmap)
		}
		s.Push(curNode)
	}
	for !s.IsEmpty() {
		popStackSetMap(s, rmap)
	}



	var head = new(common.BiTreeNode)

	for i := 0; i < len(nodes); i++ {
		curNode := nodes[i]

		//ln或者rn是作为父节点设置的
		ln := lmap[curNode]
		rn := rmap[curNode]

		if ln == nil && rn == nil {
			head = curNode
		} else if ln == nil {
			if rn.Left() == nil {
				rn.SetLeft(curNode)
			} else {
				rn.SetRight(curNode)
			}
		} else if rn == nil {
			if ln.Left() == nil {
				ln.SetLeft(curNode)
			} else {
				ln.SetRight(curNode)
			}
		} else {
			var pNode *common.BiTreeNode
			if cf(ln.Data(), rn.Data()) < 0 {
				pNode = ln
			} else {
				pNode = rn
			}

			if pNode.Left() == nil {
				pNode.SetLeft(curNode)
			} else {
				pNode.SetRight(curNode)
			}
		}

	}

	return head
}

func popStackSetMap(s *common.Stack, m map[*common.BiTreeNode]*common.BiTreeNode) {
	node := s.PopValue().(*common.BiTreeNode)
	if s.IsEmpty() {
		m[node] = nil
	} else {
		m[node] = s.Top().(*common.BiTreeNode)
	}
}

func main() {
	arr := []interface{}{3, 4, 5, 1, 2}

	relTreeNode := getMaxTreeFromArr(arr, common.Compare_int)

	l := new(common.List)



	//程序遍历数
	relTreeNode.LevelorderNo(l)
	//l.Traverse(common.Tf)
}
