package common


//单链表

type ListElm struct {
	value interface{}
	next  *ListElm
}

func (e *ListElm) GetValue() interface{} {
	return e.value
}

type List struct {
	size int
	head *ListElm
	tail *ListElm
}

func (l *List) Init() {
	l.size = 0
}

//在ele后插入一个新的元素
func (l *List) Ins_next(ele *ListElm, value interface{}) {
	new_element := &ListElm{
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
		l.head = new_element
	} else {
		if ele.next == nil {
			l.tail = new_element
		}
		new_element.next = ele.next
		ele.next = new_element
	}
	l.size ++
	return
}

//在ele元素后删除一个元素
func (l *List) Rem_next(ele *ListElm) {
	if l.size == 0 {
		return
	}

	//remove list head
	if ele == nil {
		if l.size == 1 {
			l.tail = nil
		}
		//头节点后移
		l.head = l.head.next
	} else {
		if ele.next == nil {
			return
		}
		ele.next = ele.next.next

		//移除最后一个节点
		if ele.next == nil {
			l.tail = ele
		}
	}
	l.size --
	return
}

//删除ele之前的元素
func (l *List) Rem_before(ele *ListElm) {
	if l.size == 0 {
		return
	}
	var pre *ListElm
	for m := l.head; m != nil; m = m.next {
		if m == ele {
			l.Rem_next(pre)
			return
		}
		pre = m
	}
	return
}

//遍历执行 f
func (l *List) Traverse(f func(e *ListElm, args ...interface{}) bool, args ...interface{}) {

	if l.size == 0 {
		return
	}

	for m := l.head; m != nil; m = m.next {
		if f(m, args) {
			return
		}
	}
	return
}
