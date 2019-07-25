package common

import (
	rand2 "math/rand"
	"fmt"
)

/*
快速排序 解决一般问题的最佳排序算法 分治算法
 分 ： 设定一个分割值将数据分成两部分
 治 ： 分别在两部分用递归的方式继续使用快速排序算法
 合 ： 对分割部分排序直至完成

//复杂度 O(n lgn) 不占用额外的存储空间，只使用无序数组本身的存储空间即可
*/

func Compare_int(v1, v2 interface{}) int {
	if v1.(int) > v2.(int) {
		return 1
	} else if v1.(int) < v2.(int) {
		return -1
	} else {
		return 0
	}
}

func rand() int {
	return rand2.Intn(2147483647)
}

//i,k 初试值设置为 0 size-1 分区
func Partition(data []interface{}, i, k int, cf CF) int {
	fmt.Println("DONE")
	if i == k {
		return i
	}
	var pval interface{}

	var r = make([]interface{}, 3, 3)

	r[0] = (rand() % (k - i + 1)) + i
	r[1] = (rand() % (k - i + 1)) + i
	r[2] = (rand() % (k - i + 1)) + i

	Insort(r, 3, Compare_int)

	pval = data[r[1].(int)] //取中位数的方法取出一个值作为哨兵

	for {
		//从右端开始向左搜索找到第一个小于或者等于哨兵的值
		for (cf(data[k], pval) > 0) {
			k--
		}

		//从左端开始向右搜索找到第一个大于或者哨兵的值
		for (cf(data[i], pval) < 0) {
			i++
		}

		//一直找到重合 ，则i左边均小于哨兵，右边均大于哨兵
		if (i >= k) {
			break
		} else {
			if (cf(data[i], data[k]) != 0) {
				data[i], data[k] = data[k], data[i]
			} else {
				//防止陷入死循环
				del := k - i
				if del == 1 {
					return i
				} else {
					pval = data[i+1]
				}
			}
		}
	}
	return i
}

//成功返回0，失败返回-1；分区方法调用，确定左边递归调用，右边一直确定新的哨兵，知道i的值移到最后
func Qksort1(data []interface{}, i, k int, cf CF) int {
	var j int
	for (i < k) {
		if j = Partition(data, i, k, cf); j < 0 {
			return -1
		}

		if Qksort1(data, i, j, cf) < 0 {
			return -1
		}
		i = j + 1
	}
	return 0
}

//成功返回0，失败返回-1；分区方法调用 
func Qksort2(data []interface{}, i, k int, cf CF) int {
	if i >= k {
		return 0
	}
	j := Partition(data, i, k, cf)

	if Qksort2(data, i, j, cf) < 0 {
		return -1
	}

	if Qksort2(data, j+1, k, cf) < 0 {
		return -1
	}

	return 0
}
