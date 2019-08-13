package common

/*
二叉搜索树

(1)  某个节点的左子树中的所有节点的值都比这个节点的值小

(2)  某个节点的右子树中的所有节点的值都比这个节点的值大

查找： 从根结点向下查询，若是结点值比目标值小，顺着左子树继续查找，反之顺着右子树继续查找。

插入：也是按照规则，从根结点开始，一直查询到空位置进行插入

复杂度 o(lg n) : n代表树中结点的个数 （平衡二叉树的情况）

*/

import (
	"fmt"
)

type TreeNode struct {
	data   interface{}
	rc     int //副本数，某个位置有rc个相同副本
	lchild *TreeNode
	rchild *TreeNode
	parent *TreeNode
}

type BiSearchTree struct {
	cf     CF
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
	size   int
}

func (bst *BiSearchTree) Init(args ...interface{}) {
	bst.size = 0
	if len(args) == 1 {
		bst.cf = args[0].(func(key1, key2 interface{}) int)
	}
	return
}

//插入
func (bst *BiSearchTree) Add(data interface{}) {
	bst.create = new(TreeNode)
	bst.create.data = data

	if !bst.IsEmpty() {
		bst.cur = bst.root
		for {
			if bst.cf(data, bst.cur.data) < 0 {
				//如果要插入的值比当前节点的值小，则当前节点指向当前节点的左孩子，如果
				//左孩子为空，就在这个左孩子上插入新值
				if bst.cur.lchild == nil {
					bst.cur.lchild = bst.create
					bst.create.parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.lchild
				}

			} else if bst.cf(data, bst.cur.data) > 0 {
				//如果要插入的值比当前节点的值大，则当前节点指向当前节点的右孩子，如果
				//右孩子为空，就在这个右孩子上插入新值
				if bst.cur.rchild == nil {
					bst.cur.rchild = bst.create
					bst.create.parent = bst.cur
					break
				} else {
					bst.cur = bst.cur.rchild
				}

			} else {
				//如果要插入的值在树中已经存在，副本数加一
				bst.cur.rc++
				return
			}
		}

	} else {
		bst.root = bst.create
		bst.root.parent = nil
		bst.size++
	}
}

//查找
func (bst *BiSearchTree) Search(data interface{}) *TreeNode {
	//和Add操作类似，只要按照比当前节点小就往左孩子上拐，比当前节点大就往右孩子上拐的思路
	//一路找下去，知道找到要查找的值返回即可
	bst.cur = bst.root
	for {
		if bst.cur == nil {
			return nil
		}

		if cm := bst.cf(data, bst.cur.data); cm < 0 {
			bst.cur = bst.cur.lchild
		} else if cm > 0 {
			bst.cur = bst.cur.rchild
		} else {
			if bst.cur.rc < 0 {
				return nil
			}
			return bst.cur
		}
	}
}

//删除节点 真删除
func (bst *BiSearchTree) DeleteNode(node *TreeNode) {
	if node == nil {
		return
	}

	if node.lchild == nil && node.rchild == nil {
		//如果要删除的节点没有孩子，直接删掉它就可以
		if node == bst.root {
			bst.root = nil
		} else {
			if node.parent.lchild == node {
				node.parent.lchild = nil
			} else {
				node.parent.rchild = nil
			}
		}

	} else if node.lchild != nil && node.rchild == nil {
		//如果要删除的节点只有左孩子或右孩子，让这个节点的父节点指向它的指针指向它的
		//孩子即可,孩子指向其父节点,要是删除的是根节点，让孩子为其根节点
		if node == bst.root {
			node.lchild.parent = nil
			bst.root = node.lchild
		} else {
			node.lchild.parent = node.parent
			if node.parent.lchild == node {
				node.parent.lchild = node.lchild
			} else {
				node.parent.rchild = node.lchild
			}
		}

	} else if node.lchild == nil && node.rchild != nil {
		if node == bst.root {
			node.rchild.parent = nil
			bst.root = node.rchild
		} else {
			node.rchild.parent = node.parent
			if node.parent.lchild == node {
				node.parent.lchild = node.rchild
			} else {
				node.parent.rchild = node.rchild
			}
		}
	} else {
		//如果要删除的节点既有左孩子又有右孩子，就把这个节点的直接后继的值赋给这个节
		//点，然后删除直接后继节点即可
		successor := bst.GetSuccessor(node)
		node.data = successor.data
		bst.DeleteNode(successor)
	}
}

//删除
func (bst *BiSearchTree) Delete(data interface{}) {

	node := bst.Search(data)

	bst.DeleteNode(node)
}

//伪删除
func (bst *BiSearchTree) FakeDelete(data interface{}) {

	node := bst.Search(data)

	node.rc--
}

func (bst BiSearchTree) GetRoot() *TreeNode {
	if bst.root != nil {
		return bst.root
	}
	return nil
}

func (bst *BiSearchTree) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

//中序遍历
func (bst *BiSearchTree) InOrderTravel() {
	var inOrderTravel func(node *TreeNode)

	inOrderTravel = func(node *TreeNode) {
		if node != nil {
			inOrderTravel(node.lchild)
			fmt.Printf("%v ", node.data)
			inOrderTravel(node.rchild)
		}
	}

	inOrderTravel(bst.root)
}

//递归求深度
func (bst *BiSearchTree) GetDeepth() int {
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.lchild == nil && node.rchild == nil {
			return 1
		}
		var (
			ldeepth int = getDeepth(node.lchild)
			rdeepth int = getDeepth(node.rchild)
		)
		if ldeepth > rdeepth {
			return ldeepth + 1
		} else {
			return rdeepth + 1
		}
	}

	return getDeepth(bst.root)
}

func (bst *BiSearchTree) GetMin() interface{} {
	//根据二叉查找树的性质，树中最左边的节点就是值最小的节点
	if bst.root == nil {
		return nil
	}
	bst.cur = bst.root
	for {
		if bst.cur.lchild != nil {
			bst.cur = bst.cur.lchild
		} else {
			return bst.cur.data
		}
	}
}

func (bst *BiSearchTree) GetMax() interface{} {
	//根据二叉查找树的性质，树中最右边的节点就是值最大的节点
	if bst.root == nil {
		return nil
	}
	bst.cur = bst.root
	for {
		if bst.cur.rchild != nil {
			bst.cur = bst.cur.rchild
		} else {
			return bst.cur.data
		}
	}
}

//前驱节点: 按小的方向找：找到可以替换node的节点，值替换到到node后，不影响二叉搜索树的结构
func (bst *BiSearchTree) GetPredecessor(node *TreeNode) *TreeNode {
	getMax := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.rchild != nil {
				node = node.rchild
			} else {
				return node
			}
		}
	}

	if node != nil {
		if node.lchild != nil {
			//如果这个节点有左孩子，那么它的直接前驱就是它左子树的最右边的节点，因为比这
			//个节点值小的节点都在左子树，而这些节点中值最大的就是这个最右边的节点
			return getMax(node.lchild)
		} else {
			//如果这个节点没有左孩子，那么就沿着它的父节点找，知道某个父节点的父节点的右
			//孩子就是这个父节点，那么这个父节点的父节点就是直接前驱
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.rchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

//后继节点：按大的的找：找到可以替换node的节点，值替换到到node后，不影响二叉搜索树的结构
func (bst *BiSearchTree) GetSuccessor(node *TreeNode) *TreeNode {

	//一直找到子节点中左节点不存在的节点
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.lchild != nil {
				node = node.lchild
			} else {
				return node
			}
		}
	}

	if node != nil {
		//右子树不为空可以采用，向下找，
		if node.rchild != nil {
			return getMin(node.rchild)
		} else { //向上找，找到第一个第一个满足是左节点的节点，返回其父节点
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.lchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (bst *BiSearchTree) Clear() {
	bst.root = nil
	bst.cur = nil
	bst.create = nil
}
