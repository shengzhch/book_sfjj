package common

//单链表
import (
	"fmt"
)

type ListElm struct {
	value interface{}
	next  *ListElm
}

func (e *ListElm) GetValue() interface{} {
	return e.value
}

func (e *ListElm) Next() *ListElm {
	return e.next
}

func (e *ListElm) SetValue(value interface{}) {
	e.value = value
}

func (e *ListElm) SetNext(next *ListElm) {
	e.next = next
}

//包含头尾节点，操作时要维护信息
type List struct {
	size int
	head *ListElm
	tail *ListElm
}

func (l *List) Init() {
	l.size = 0
}

func (l *List) Head() *ListElm {
	return l.head
}

func (l *List) Tail() *ListElm {
	return l.tail
}

func (l *List) SetHead(head *ListElm) {
	l.head = head
}

func (l *List) SetTail(tail *ListElm) {
	l.tail = tail
}

func (l *List) Size() int {
	return l.size
}

//在ele后插入一个新的元素,ele 为nil时表示在插入的节点作为头节点
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
	l.size++
	return
}

//为什么经常采用删除某个节点后续节点的操作，而不是直接删除某个节点

//删除节点时需要确定前一个节点的信息，直接删除某个节点也是利用了前一个节点的信息，直接删除获取有重复操作
//所以一般采用删除节点后续节点的方法删除某个节点

//在ele元素后删除一个元素 ，若ele为nil，表示移除头节点
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

		//移除的是最后一个节点时需要重置尾节点
		if ele.next == nil {
			l.tail = ele
		}
	}
	l.size--
	return
}

//删除ele元素。
func (l *List) Rem_ele(ele *ListElm) {
	if l.size == 0 || ele == nil {
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

//删除ele的元素。
func (l *List) Rem_ele2(ele *ListElm) {
	if l.size == 0 || ele == nil {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}

	var pre *ListElm
	//需要记录前一个节点的信息
	for m := l.head; m != nil; m = m.next {
		if m == ele {
			if m == l.head {
				l.head = l.head.next
			} else if m == l.tail {
				l.tail = pre
			} else {
				pre.next = m.next
			}
			l.size--
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
		if m == l.tail {
			return
		}
	}
	return
}

//根据节点关系遍历遍历执行 f
func (l *ListElm) Traverse(f func(e *ListElm, args ...interface{}) bool, args ...interface{}) {

	for m := l; m != nil; m = m.next {
		if f(m, args) {
			return
		}
	}
	return
}

//定义一个打印的函数
var (
	Tf = func(e *ListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}
)
