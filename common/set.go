package common

//集合 --- 插入 删除 并 交 差 属于 包含子集 相等 大小
type Set List

func (s *Set) Insert(data interface{}) {
	if !s.Belong(data) {
		List(*s).Ins_next(s.tail, data)
	}
}

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

	List(*s).Rem_next(prev)

	return
}

func (s *Set) Union(s2 *Set) *Set {

	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		List(*rel).Ins_next(rel.tail, m.value)
	}
	for m := s2.head; m != nil; m = m.next {
		if rel.Belong(m.value) {
			continue
		} else {
			List(*rel).Ins_next(rel.tail, m.value)
		}
	}
	return rel
}

func (s *Set) Intersection(s2 *Set) *Set {
	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		if s2.Belong(m.value) {
			List(*rel).Ins_next(rel.tail, m.value)
		}
	}
	return rel

}

func (s *Set) Difference(s2 *Set) *Set {
	var rel = new(Set)

	for m := s.head; m != nil; m = m.next {
		if !s2.Belong(m.value) {
			List(*rel).Ins_next(rel.tail, m.value)
		}
	}
	return rel
}

func (s *Set) Belong(data interface{}) bool {

	for m := s.head; m != nil; m = m.next {
		if m.value == data {
			return true
		}
	}
	return false
}

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

func (s *Set) Size() int {
	return List(*s).size
}
