package common

/*
哈希 哈希表包含一个数组，支持通过key通过哈希运算获得下标来访问数组中的元素。

与各类的key相比，哈希表的条目数相应较少，因此，绝大多数哈希函数会将不同的key映射到表中相同的槽位上，称为哈希冲突
*/

/* 链式哈希表 根本是一组链表，每组链表称为一个桶，通过key的哈希运算确定属于哪个桶。目标是尽可能的是元素均匀的分布在每个桶上。
哈希表的负载因子 a = n/m 。n为表中元素的个数，m为桶的个数。一般情况下 我们往往会检索元素个数大于负载因子的桶。

选择哈希函数 h(k) = x . x为整型，表示地址

取余法

h(k) = k mod m ; k为键值，m为槽位(桶)的个数

乘法
k乘以常数A取结果的小数部分，然后乘以m取结果的整数部分。A常取 0.618。

h(k) = FLOOR(（k*A mod 1）* m )
*/

//能够较好处理字符串的哈希函数，通过一系列的位操作将键强制转化为整数。所有的这些整数都是通过取余法得到的
func Hashpjw(key string, tableSize int) int {
	var val = 0
	for i := 0; i < len(key) && key[i] != byte(0); i++ {
		var tmp int
		val = (val << 4) + int(byte(key[i]))
		if tmp == (val & 0xf0000000) {
			val = val ^ (tmp >> 24)
			val = val ^ tmp
		}
	}
	return val % tableSize
}

//链式哈希表 维护一个链表数组 取余法
type Chtable struct {
	buckets int
	size    int
	table   []List

	//f(key)%buckets 确定槽位 辅助哈希函数
	f func(key interface{}) int

	//
	match func(key1, key2 interface{}) bool
}

//初始化
func (l *Chtable) Init(b int, f func(key interface{}) int, match func(key1, key2 interface{}) bool) {
	l.table = make([]List, 0, b)
	l.buckets = b
	l.size = 0
	l.f = f
	l.match = match
}

//是否存在 根据match进行匹配
func (l *Chtable) Look_up(data interface{}) bool {
	buc := l.f(data) % l.buckets

	f := func(e *ListElm, args ...interface{}) bool {
		if len(args) != 2 {
			return true
		}

		if l.match(e.value, args[0]) {
			*args[1].(*bool) = true
			return true
		}
		return false
	}

	var rel = new(bool)
	l.table[buc].Traverse(f, data, rel)

	return *rel
}

//插入 插入的元素也是data
func (l *Chtable) Insert(data interface{}) {
	if l.Look_up(data) {
		return
	}

	buc := l.f(data) % l.buckets

	l.table[buc].Ins_next(nil, data)
	l.size++
	return
}

//删除
func (l *Chtable) Remove(data interface{}) {
	buc := l.f(data) % l.buckets
	bucket := l.table[buc]
	var pre *ListElm
	for m := bucket.head; m != nil; m = m.next {
		if l.match(m.value, data) {
			bucket.Rem_next(pre)
			l.size--
			return
		}
		pre = m
	}

	return
}

//不好，只是为了验证Traverse 有重复操作
func (l *Chtable) Remove_NoGood(data interface{}) {
	buc := l.f(data) % l.buckets

	f := func(e *ListElm, args ...interface{}) bool {
		if len(args) != 3 {
			return true
		}

		if l.match(e.value, args[0]) {
			*args[1].(*bool) = true
			*args[2].(*ListElm) = *e
			return true
		}
		return false
	}

	var rel = new(bool)
	var finde = new(ListElm)
	l.table[buc].Traverse(f, data, rel, finde)

	if *rel {
		l.table[buc].Rem_ele(finde)
		l.size--
	}

	return
}
