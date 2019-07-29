/*
计数排序是一种高效的线性排序，通过计算一个集合中的元素出现的次数来确定集合如何排列，只能排列用于整型或者可以用整型来表示的数据集合
*/
package common


//扩展了原来书中所给的功能，满足取最小值为负数的情况排序
//min 为最大值 ，min为最小值
func Ctsort(data []int, size, min, max int) int {
	var count = make([]int, max-min+1) //取值范围区间

	var tmp = make([]int, size)
	var i, j int

	//min-max 映射到 0 max -min
	for i = 0; i <= max-min; i++ {
		count[i] = 0
	}

	for j = 0; j < size; j++ {
		offset := data[j] - min
		count[offset] = count[offset] + 1
	}

	//每个位置加上偏移量
	for i = 0; i <= max-min; i++ {
		count[i] = count[i] + count[i-1]
	}

	for j = 0; j < size; j++ {
		index := count[data[j]-min] - 1
		tmp[index] = data[j]
		count[data[j]-min] = count[data[j]-min] - 1
	}

	copy(data[:], tmp[:])
	return 0
}
