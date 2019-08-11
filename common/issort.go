package common

//插入排序 复杂度 O(n^2),不占用额外的存储空间，只使用数组本身的存储空间即可

type CF func(key1, key2 interface{}) int

//成功返回0，返回-1 ；
func Insort(data []interface{}, size int, cf CF) int {
	var i, j int
	var key interface{}
	for j = 1; j < size; j++ {
		key = data[j]

		i = j - 1

		//一步一步把key移动到合适的位置，前面的数据是已经排好序的, 用>而不用等于是为了保持算法的稳定性
		for i >= 0 && cf(data[i], key) > 0 {
			data[i+1], data[i] = data[i], data[i+1]
			i--
		}
	}
	return 0
}
