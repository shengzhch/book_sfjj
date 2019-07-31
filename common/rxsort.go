/*
基数排序
数据按位分开，从最低有效位到最高有效位进行比较，依次排序 O(pn + pk)
*/
package common

import "math"

//p 为位的个数 k为基数
func Rxsort(data []int, size, p, k int) int {
	var (
		count                = make([]int, k)
		tmp                  = make([]int, size)
		
		index, pval, i, j, n int
	)

	//对每一位上运行计数排序
	for n = 0; n < p; n++ {

		for i = 0; i < k; i++ {
			count[i] = 0
		}

		pval = int(math.Pow(float64(k), float64(n)))

		for j = 0; j < size; j++ {
			index = (int)(data[j]/pval) % k
			count[index] = count[index] + 1
		}
		
		
		//偏移量
		for i = 1; i < k; i++ {
			count[i] = count[i] + count[i-1]
		}

		for j = 0; j < size; j++ {
			index = (int)(data[j]/pval) % k
			tmp[count[index]-1] = data[j]
			count[index] = count[index] - 1
		}

		copy(data[:], tmp[:])
	}

	return 0
}
