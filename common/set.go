package common

//集合 --- 无序且独立
//同样用链表做底层的数据
type Set List

//属于
func (s *Set) Belong(data interface{}) bool {

	for m := s.head; m != nil; m = m.next {
		if m.value == data {
			return true
		}
	}
	return false
}

//插入
func (s *Set) Insert(data interface{}) {
	if !s.Belong(data) {
		(*List)(s).Ins_next(s.tail, data)
	}
}


//删除
func (s *Set) Remove(data interface{}) {
	var prev *ListElm
	var find *ListElm

	for m := s.head; m != nil; m = m.next {
		if m.value == data {
			find = m
			break
		}
		prev = m
	}
	if find == nil {
		return
	}

	(*List)(s).Rem_next(prev)

	return
}

//并集 返回一个新的集合
func (s *Set) Union(s2 *Set) *Set {

	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		(*List)(s).Ins_next(rel.tail, m.value)
	}
	for m := s2.head; m != nil; m = m.next {
		if rel.Belong(m.value) {
			continue
		} else {
			(*List)(s).Ins_next(rel.tail, m.value)
		}
	}
	return rel
}

//交集 返回一个新的集合
func (s *Set) Intersection(s2 *Set) *Set {
	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		if s2.Belong(m.value) {
			(*List)(s).Ins_next(rel.tail, m.value)
		}
	}
	return rel

}

//差集 返回一个新的集合
func (s *Set) Difference(s2 *Set) *Set {
	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		if !s2.Belong(m.value) {
			(*List)(s).Ins_next(rel.tail, m.value)
		}
	}
	return rel
}


//是否包含
func (s *Set) Contain(s2 *Set) bool {
	if s2.Size() > s.Size() {
		return false
	}
	for m := s2.head; m != nil; m = m.next {
		if !s.Belong(m.value) {
			return false
		}
	}
	return true
}


//是否相等
func (s *Set) Equal(s2 *Set) bool {
	if s.Size() != s2.Size() {
		return false
	}
	for m := s2.head; m != nil; m = m.next {
		if !s.Belong(m.value) {
			return false
		}
	}
	return true
}

//集合元素个数
func (s *Set) Size() int {
	return (*List)(s).size
}
