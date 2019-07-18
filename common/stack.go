package common

//栈先先出。栈顶进出
//用队列表示栈， 队列头表示栈顶
type Stack List


func (s *Stack) Push(data interface{}) {
	List(*s).Ins_next(nil, data)
}

func (s *Stack) Pop(data interface{}) {
	List(*s).Rem_next(nil)
}

func (s *Stack) Top() interface{} {
	return s.head.value
}
