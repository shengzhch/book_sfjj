package common

//队列 先进先出 对尾插入 对头删除 也用单链表表示队列
type Queue List

//对尾插入
func (s *Queue) EnQueue(data interface{}) {
	(*List)(s).Ins_next(s.tail, data)
}

//对头删除
func (s *Queue) DeQueue() {
	(*List)(s).Rem_next(nil)
}

//对头删除
func (s *Queue) DeQueueWithValue() interface{} {
	rel := List(*s).head
	(*List)(s).Rem_next(nil)
	return rel.value
}

//对头删除
func (s *Queue) Size() int {
	return s.size
}

//对头
func (s *Queue) Head() interface{} {
	return s.head.value
}
