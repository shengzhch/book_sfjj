//优先队列 ：对尾插入，按优先级大小出队列
//最大优先队列 无论入队顺序，优先级最高的元素优先出队
//最小优先队列 无论入队顺序，优先级最低的元素优先出队

package common

type Pqueue Heap

func (s *Pqueue) Init(args ...interface{}) {
	(Heap)(*s).Init(args)
}

func (s *Pqueue) Insert(data interface{}) {
	(Heap)(*s).Insert(data)
}
func (s *Pqueue) Extract(args ...interface{}) {
	(Heap)(*s).Extract()
}
func (s *Pqueue) Size() int {
	return (Heap)(*s).size
}

//优先级最高的元素
func (s *Pqueue) Peek() interface{} {
	return (Heap)(*s).tree[0]
}
