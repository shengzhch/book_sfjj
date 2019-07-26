package common

import (
	"errors"
)

type BiTreeNode struct {
	data  interface{}
	left  *BiTreeNode
	right *BiTreeNode
}

type BiTree struct {
	size    int
	compare func(key1, key2 interface{}) int
	root    *BiTreeNode
}

var (
	Wrong = errors.New("Wrong")
)

func (bs *BiTree) Init(args ...interface{}) {
	bs.size = 0
	if len(args) == 1 {
		bs.compare = args[0].(func(key1, key2 interface{}) int)
	}
	return
}

//将新节点插入到二叉树中且node节点的左子结点，新节点的数据为data。若node为nil且二叉树为空树，则新结点作为根结点
func (bs *BiTree) Ins_left(node *BiTreeNode,
	data interface{}) (e error) {

	var pos = &BiTreeNode{data: data}
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.left != nil {
			return Wrong
		}
		node.left = pos
	}

	bs.size++
	return
}

func (bs *BiTree) Ins_right(node *BiTreeNode, data interface{}) (e error) {
	var pos = &BiTreeNode{data: data}
	if node == nil {
		if bs.size > 0 {
			return Wrong
		}

		bs.root = pos
	} else {
		if node.right != nil {
			return Wrong
		}
		node.right = pos
	}

	bs.size++
	return nil
}

//将指定结点node的左子树移除,node 为nil，以根结点为起始结点
func (bs *BiTree) Rem_left(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.left
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.Rem_left(*pos)
		bs.Rem_right(*pos)
		*pos = nil
		bs.size --
	}
	return
}

func (bs *BiTree) Rem_right(node *BiTreeNode) {
	if bs.size == 0 {
		return
	}

	var pos **BiTreeNode

	if node == nil {
		pos = &bs.root
	} else {
		pos = &node.right
	}

	//后序遍历删除所有的结点
	if *pos != nil {
		bs.Rem_left(*pos)
		bs.Rem_right(*pos)
		*pos = nil
		bs.size --
	}
	return
}

//将两棵二叉树合并为单棵二叉树
func Merge(bs, bs2 *BiTree, data interface{}) (*BiTree, error) {
	rel := &BiTree{size: 0}
	if e := rel.Ins_left(nil, data); e != nil {
		return nil, e
	}
	rel.root.left = bs.root
	rel.root.right = bs2.root
	rel.size = rel.size + bs.size + bs2.size

	bs.size = 0
	bs.root = nil
	bs2.size = 0
	bs2.root = nil
	return rel, nil
}

func (bs *BiTree) Size() int {
	return bs.size
}

func (bs *BiTree) Root() *BiTreeNode {
	return bs.root
}

func (node *BiTreeNode) Preorder(l *List) {

	if !node.Is_eob() {
		l.Ins_next(l.tail, node.data)

		if !node.left.Is_eob() {
			node.left.Preorder(l)
		}

		if !node.right.Is_eob() {
			node.right.Preorder(l)
		}
	}
}

func (node *BiTreeNode) Inorder(l *List) {
	if !node.Is_eob() {

		if !node.left.Is_eob() {
			node.left.Inorder(l)
		}

		l.Ins_next(l.tail, node.data)

		if !node.right.Is_eob() {
			node.right.Inorder(l)
		}
	}
}

func (node *BiTreeNode) Postorder(l *List) {
	if !node.Is_eob() {
		if !node.left.Is_eob() {
			node.left.Postorder(l)
		}

		if !node.right.Is_eob() {
			node.right.Postorder(l)
		}

		l.Ins_next(l.tail, node.data)
	}
}

//非递归遍历 遇到一个结点，插入到链表中，然后压栈，然后遍历左子树，左子树遍历完了，出栈，指向右结点，继续做
func (node *BiTreeNode) PreorderNo(l *List) {
	s := new(Stack)
	s.Init()
	var tmp = new(BiTreeNode)

	tmp = node

	for (tmp != nil || s.Size() > 0) {
		for (tmp != nil) {
			l.Ins_next(l.tail, tmp.data)
			s.Push(tmp)
			tmp = tmp.left
		}
		if s.Size() > 0 {
			tmp = s.PopValue().(*BiTreeNode)
			tmp = tmp.right
		}
	}
}

//非递归遍历 遇到一个结点，压栈，一直找到左结点为空，出栈，插入，指向右结点，继续做
func (node *BiTreeNode) InorderrNo(l *List) {
	s := new(Stack)
	s.Init()
	var tmp = new(BiTreeNode)

	tmp = node

	for (tmp != nil || s.Size() > 0) {
		//一路向左压栈
		for (tmp != nil) {
			s.Push(tmp)
			tmp = tmp.left
		}

		if s.Size() > 0 {
			tmp = s.PopValue().(*BiTreeNode)
			l.Ins_next(l.tail, tmp.data) //插入
			tmp = tmp.right
		}
	}
}

//两个栈来做辅助存储 依次把根结点 右儿子 左儿子压栈，出栈的顺序就是后序遍历，处理过程于前序遍历相仿
func (node *BiTreeNode) PostorderNo(l *List) {
	s1 := new(Stack)
	s1.Init()
	s2 := new(Stack)
	s2.Init()

	var tmp = new(BiTreeNode)

	tmp = node

	for (tmp != nil || s1.Size() > 0) {
		//一路向右压栈
		for (tmp != nil) {
			s1.Push(tmp)
			s2.Push(tmp)
			tmp = tmp.right
		}

		if s1.Size() > 0 {
			tmp = s1.PopValue().(*BiTreeNode)
			tmp = tmp.left
		}
	}

	for (s2.Size() > 0) {
		tmp = s1.PopValue().(*BiTreeNode)
		l.Ins_next(l.tail, tmp.data) //插入
	}
}

//层级遍历 根结点入队列，然后执行循环，出队列，左儿子入队，右儿子入队，然后继续做
func (node *BiTreeNode) LevelorderNo(l *List) {
	q := new(Queue)
	var tmp = new(BiTreeNode)

	tmp = node

	if tmp == nil {
		return
	}
	q.EnQueue(tmp)

	for (q.Size() > 0) {
		tmp = q.DeQueueWithValue().(*BiTreeNode)
		l.Ins_next(l.tail, tmp.data) //插入
		if tmp.left != nil {
			q.EnQueue(tmp.left)
		}
		if tmp.right != nil {
			q.EnQueue(tmp.right)
		}
	}
}

//判断node是不是二叉树的叶子结点
func (node *BiTreeNode) Is_leaf() bool {
	return node.left == nil && node.right == nil
}

//判断node的终止 end of bitree
func (node *BiTreeNode) Is_eob() bool {
	return node == nil
}

//
func (node *BiTreeNode) Right() *BiTreeNode {
	return node.right
}

//
func (node *BiTreeNode) Left() *BiTreeNode {
	return node.left
}

//
func (node *BiTreeNode) Data() interface{} {
	return node.data
}
