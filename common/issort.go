package common

//插入排序 复杂度 O(n^2),不占用额外的存储空间，只使用数组本身的存储空间即可

type CF func(key1, key2 interface{}) bool

//成功返回0，返回-1 ；
func Insort(data []interface{}, size int, cf CF) int {
	var i, j int
	var key interface{}
	for j = 1; j < size; j++ {
		key = data[j]

		i = j - 1

		for i >= 0 && cf(data[i], key) {
			data[i+1], data[i] = data[i], data[i+1]
			i--
		}
	}
	return 0
}
