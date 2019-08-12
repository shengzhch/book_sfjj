package common

//循环链表 底层是单链表
type CList List

//在ele后插入一个新的元素
func (l *CList) Ins_next(ele *ListElm, value interface{}) {
	new_element := &ListElm{
		value: value,
	}

	if l.size == 0 {
		new_element.next = new_element
		l.head = new_element
		l.tail = new_element
	} else { //忽略了头节点插入的情况 即ele不能为nil
		new_element.next = ele.next
		ele.next = new_element
	}

	//重新设置tail
	if ele == l.tail {
		l.tail = new_element
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
		l.tail = nil
	} else {
		ele.next = ele.next.next
		next := ele.next
		//删除节点是头结点，需要更新头结点
		if next == l.head {
			l.head = next.next
		}

		if next == l.tail {
			l.tail = ele
		}
	}
	l.size --
	return
}
