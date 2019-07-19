package common

//链表
type CList struct {
	size int
	head *ListElm
}

func (l *CList) Init() {
	l.size = 0
}

//在ele后插入一个新的元素
func (l *CList) Ins_next(ele *ListElm, value interface{}) {
	new_element := &ListElm{
		value: value,
	}

	if l.size == 0 {
		new_element.next = new_element
		l.head = new_element
	} else {
		new_element.next = ele.next
		ele.next = new_element
	}

	l.size ++
	return
}

//在ele元素后删除一个元素
func (l *CList) Rem_next(ele *ListElm) {
	if l.size == 0 {
		return
	}

	if ele.next == ele {
		l.head = nil
	} else {
		next := ele.next
		ele.next = ele.next.next

		//删除节点是头结点，需要更新头结点
		if next == l.head {
			l.head = next.next
		}
	}
	l.size --
	return
}