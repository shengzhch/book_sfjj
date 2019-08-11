package common

//栈 后进先出。栈顶进行插入删除。 用链表表示，表头表示栈顶
type Stack List

func (s *Stack) Init() {
	(*List)(s).Init()
}

func (s *Stack) Push(data interface{}) {
	(*List)(s).Ins_next(nil, data)
}

func (s *Stack) Pop() {
	(*List)(s).Rem_next(nil)
}

func (s *Stack) PopValue() interface{} {
	rel := s.head.value
	(*List)(s).Rem_next(nil)
	return rel
}

func (s *Stack) Top() interface{} {
	return s.head.value
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
