package common

import (
	"fmt"
)

//双向链表 每个节点包含指向前后方节点的指针
type DListElm struct {
	value interface{}
	pre   *DListElm
	next  *DListElm
}

func (e *DListElm) GetValue() interface{} {
	return e.value
}

func (e *DListElm) Next() *DListElm {
	return e.next
}

func (e *DListElm) Pre() *DListElm {
	return e.pre
}

func (e *DListElm) SetValue(value interface{}) {
	e.value = value
}

func (e *DListElm) SetNext(next *DListElm) {
	e.next = next
}

func (e *DListElm) SetPre(pre *DListElm) {
	e.pre = pre
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

		//在链表尾部插入，更新尾节点
		if ele.next == nil {
			l.tail = new_element
		} else {
			//原ele的后续节点指向新节点
			ele.next.pre = new_element
		}

		//
		ele.next = new_element
	}
	l.size ++
	return
}

//严格模式 在ele后插入一个新的元素
func (l *DList) Ins_next_strict(ele *DListElm, value interface{}) {
	new_element := &DListElm{
		value: value,
	}

	//在l有元素的情况下必须制定要在哪一个节点之后进行插入，nil直接返回
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
			//头节点前置指针置为空
			l.head.pre = nil
		}
	} else {
		//ele的前一个节点的后置指针指向ele的后一个节点
		ele.pre.next = ele.next

		//后续节点为空，删除的是尾节点。需要跟新尾节点
		if ele.next == nil {
			l.tail = ele.pre
		} else {
			//ele的后一个节点的前置指针指向ele的前一个节点
			ele.next.pre = ele.pre
		}
	}

	l.size--
}

//遍历执行 f
func (l *DList) Traverse(f func(e *DListElm, args ...interface{}) bool, args ...interface{}) {

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

//倒叙遍历执行 f
func (l *DList) ReTraverse(f func(e *DListElm, args ...interface{}) bool, args ...interface{}) {

	if l.size == 0 {
		return
	}

	for m := l.tail; m != nil; m = m.pre {
		if f(m, args) {
			return
		}
	}
	return
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
func (l *DList) SetHead(head *DListElm) {
	l.head = head
}

//
func (l *DList) GetTail() *DListElm {
	return l.tail
}

//
func (l *DList) SetTail(tail *DListElm) {
	l.tail = tail
}


var (
	Tdf = func(e *DListElm, args ...interface{}) bool {
		if e == nil {
			return true
		}
		fmt.Println(e.GetValue())
		return false
	}
)
