package common

//栈先先出。栈顶进出
//用队列表示栈， 队列头表示栈顶
type Stack List

func (s *Stack) Init() {
	List(*s).Init()
}

func (s *Stack) Push(data interface{}) {
	List(*s).Ins_next(nil, data)
}

func (s *Stack) Pop() {
	List(*s).Rem_next(nil)
}

func (s *Stack) PopValue() interface{} {
	rel := s.head.value
	List(*s).Rem_next(nil)
	return rel
}

func (s *Stack) Top() interface{} {
	return s.head.value
}

func (s *Stack) Size() int {
	return s.size
}
