/*
二分查找 不断将有序数据集进行对半分割，并检查中间的元素

*/
package common

import "fmt"

func Bisearch(data []interface{}, target interface{}, size int, cf CF) int {
	var left, mid, right int

	left = 0
	right = size - 1

	for left <= right {
		mid = (left + right) / 2
		fmt.Println("mid ",mid)
		switch cf(data[mid], target) {
		case -1:
			left = mid + 1
		case 0:
			return mid
		case 1:
			right = mid - 1

		}
	}
	return -1
}
