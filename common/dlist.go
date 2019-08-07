package common

//双向链表

type DListElm struct {
	value interface{}
	pre   *DListElm
	next  *DListElm
}

type DList struct {
	size int
	head *DListElm
	tail *DListElm
}

func (l *DList) Init() {
	l.size = 0
}

//在ele后插入一个新的元素
func (l *DList) Ins_next(ele *DListElm, value interface{}) {
	new_element := &DListElm{
		value: value,
	}

	//代表新插入的元素位于链表的头部
	if ele == nil {
		//表示为空链表，要更新链表的尾部
		if l.size == 0 {
			l.tail = new_element
		}

		//更新头部链表头部
		new_element.next = l.head
		if l.head != nil {
			l.head.pre = new_element
		}
		l.head = new_element
	} else {
		new_element.pre = ele
		new_element.next = ele.next

		//在链表尾部插入
		if ele.next == nil {
			l.tail = new_element
		} else {
			ele.next.pre = new_element
		}
		ele.next = new_element
	}
	l.size ++
	return
}

//在ele后插入一个新的元素
func (l *DList) Ins_next_strict(ele *DListElm, value interface{}) {
	new_element := &DListElm{
		value: value,
	}

	if ele == nil && l.size > 0 {
		return
	}

	if l.size == 0 {
		l.head = new_element
		l.tail = new_element
	} else {
		new_element.pre = ele
		new_element.next = ele.next

		if ele.next == nil {
			l.tail = new_element
		} else {
			ele.next.pre = new_element
		}
		ele.next = new_element
	}

	l.size ++
	return
}

//删除ele 元素
func (l *DList) Rem_ele(ele *DListElm) {
	if ele == nil || l.size == 0 {
		return
	}

	if ele == l.head {
		l.head = ele.next
		//只有一个节点
		if l.head == nil {
			l.tail = nil
		} else {
			ele.next.pre = nil
		}
	} else {

		ele.pre.next = ele.next

		if ele.next == nil {
			l.tail = ele.pre
		} else {
			ele.next.pre = ele.pre
		}
	}

	l.size--
}

//
func (l *DList) Size() int {
	return l.size
}

//
func (l *DList) GetHead() *DListElm {
	return l.head
}

//
func (l *DList) GetHeadValue() interface{} {
	return l.head.value
}

//
func (l *DList) GetTail()*DListElm{
	return l.tail
}

//
func (l *DList) GetTailValue() interface{} {
	return l.tail.value
}
