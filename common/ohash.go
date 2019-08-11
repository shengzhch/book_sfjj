package common

/*
与链式哈希表把元素存放在每个地址的桶中不同，开地址哈希表直接存放在表本身中。

开地址哈希表 ：在开地址哈希表中解决冲突的方法是探查这个表，一直找到可以放置元素的槽位.
如果要删除一个元素，则要一直找到该元素或者确定一个空槽

开地址哈希表中包含的元素 n 不可能大于表中槽位的数量  m
负载因子 小于等于 1  （n/m）

假设进行均匀散列，我们能够在一个开地址哈希表中探查的槽位个数是 1/(1-a)，即 m/（m-n）。当a 趋于1时，该值增大趋于m，探查效率降低。
对于对时间敏感的应用，可以增大哈希表的空间

h(k,i)= x ; i是到目前为止探查的次数，x是得到的哈希编码

线性探查

h(k,i) = (h'(k)+i) mod m  -----> h'(x) 为辅助哈希函数，尽可能的将元素随机和均匀的分布在表中 例如取余法.

对m没有限制，但遇到基本聚集的时候，（元素不能均匀分布在哈希表中）有可能会探测多次才能找到空槽，降低探查效率


双散列

h(k,i) = (h'(k)+ih"(k)) mod m

通常情况下 h'(k) = k mod m  ; h"(k) = 1+(k mod m') ,其中m'略小于m。m必须为2次幂或者是一个素数。

需要限定m，但探查能产生较好的元素分布
*/

var (
	vacated interface{} = "" //表示删除后腾出的空间 c语言中为一个指针，删除后不释放内存，减少内存分配
)

//开地址哈希表 维护一个数组 用双散列法实现
type Ohtable struct {
	positions int          //
	vacated   *interface{} //指向全局变量vacated,删除某个元素后，用vacated填充
	size      int
	table     []interface{}

	h1 func(key interface{}) int
	h2 func(key interface{}) int

	//
	match func(key1, key2 interface{}) bool
}

//初始化
func (l *Ohtable) Init(p int, h1, h2 func(key interface{}) int, match func(key1, key2 interface{}) bool) {
	l.table = make([]interface{}, 0, p)
	l.positions = p
	l.size = 0
	l.h1 = h1
	l.h2 = h2
	l.match = match
	l.vacated = &vacated
}

//是否存在 根据match进行匹配
func (l *Ohtable) Look_up(data interface{}) bool {

	var p, i int

	for i = 0; i < l.positions; i++ {
		p = (l.h1(data) + i*l.h2(data)) % l.positions
		if l.table[p] == nil {
			return false
		} else if l.match(l.table[p], data) {
			return true
		}
	}

	return false
}

//插入 插入的元素也是data。返回结果  -1 插入失败，1为已存在，0 为插入成功
func (l *Ohtable) Insert(data interface{}) int {
	var p, i int

	if l.size == l.positions {
		return -1
	}

	if l.Look_up(data) {
		return 1
	}

	for i = 0; i < l.positions; i++ {
		p = (l.h1(data) + i*l.h2(data)) % l.positions
		if l.table[p] == nil || l.table[p] == *l.vacated {
			l.table[p] = data
			l.size++
			return 0
		}
	}

	//一般不发生，若发生则说明哈希函数 h1或者h2 选择的不合适
	return -1
}

//删除
func (l *Ohtable) Remove(data interface{}) int {
	var p, i int

	for i = 0; i < l.positions; i++ {
		p = (l.h1(data) + i*l.h2(data)) % l.positions
		if l.table[p] == nil {
			return -1
		} else if l.table[p] == *l.vacated {
			continue
		} else if l.match(l.table[p], data) {
			l.table[p] = *l.vacated
			l.size--
			return 0
		}
	}

	//没有找到匹配的key
	return -1
}
