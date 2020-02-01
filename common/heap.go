/*
堆 完全二叉树 子节点的值比父结点的值小（大），根结点是树中最大（小）的结点，称为最大（小）值堆。
// 用数组的方式存储数据，二叉树的层级遍历，左平衡树
*/

package common

type Heap struct {
	size int
	// 大于时返回1
	compare func(key1, key2 interface{}) int
	//层次遍历保存的数组
	tree []interface{}
}

func (h *Heap) Init(args ...interface{}) {
	h.size = 0
	h.tree = *new([]interface{})
	if len(args) == 1 {
		h.compare = args[0].(func(key1, key2 interface{}) int)
	}
	return
}

var (
	//
	heap_parent = func(npos int) int {
		return (npos - 1) / 2
	}

	heap_left = func(npos int) int {
		return (npos * 2) + 1
	}

	heap_right = func(npos int) int {
		return (npos * 2) + 2
	}
)

//规则的最大堆处理
func (h *Heap) Insert(data interface{}) {
	var ipos, ppos int

	h.tree = append(h.tree, data)

	ipos = h.size
	ppos = heap_parent(ipos)

	for (ipos > 0 && h.compare(h.tree[ppos], h.tree[ipos]) < 0) {
		h.tree[ppos], h.tree[ipos] = h.tree[ipos], h.tree[ppos]
		ipos = ppos
		ppos = heap_parent(ipos)
	}
	h.size++
}

//提取 第一个元素,提取后调整数组保证堆的正确顺序
func (h *Heap) Extract() interface{} {
	if h.size == 0 {
		return nil
	}

	data := h.tree[0]

	save := h.tree[h.size-1]

	if h.size > 1 {
		h.tree = h.tree[:h.size-1]
		h.size--
	} else {
		h.tree = nil
		h.size = 0
		return data
	}

	//最后一个元素提前
	h.tree[0] = save
	mpos := 0
	ipos := 0
	lpos := heap_left(ipos)
	rpos := heap_right(ipos)

	for {
		lpos = heap_left(ipos)
		rpos = heap_right(ipos)

		if lpos < h.size && h.compare(h.tree[lpos], h.tree[ipos]) > 0 {
			mpos = lpos
		} else {
			mpos = ipos
		}
		if rpos < h.size && h.compare(h.tree[rpos], h.tree[mpos]) > 0 {
			mpos = rpos
		}

		if (mpos == ipos) {
			break
		} else {
			h.tree[mpos], h.tree[ipos] = h.tree[ipos], h.tree[mpos]
			ipos = mpos
		}

	}

	return data
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Top() interface{} {
	if h.size > 0 {
		return h.tree[0]
	}
	return nil
}

func (h *Heap) Data() []interface{} {
	return h.tree
}
