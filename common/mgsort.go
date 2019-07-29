/*
归并排序 ：将一个无序元素集合分割成许多个只包含一个元素的集，然后不断的将这些小集合合并，需要额外的存储空间
时间复杂度 O(nlg n)
*/

package common

//合并到一起，j是分割的位置
func merge(data []interface{}, i, j, k int, cf CF) int {
	var m = make([]interface{}, k-i+1)
	ipos := i
	jpos := j + 1
	mpos := 0

	for (ipos <= j || jpos <= k) {
		//处理两组数据剩余的部分,肯定只有一组剩余
		if (ipos > j) {
			for (jpos <= k) {
				m[mpos] = data[jpos]
				jpos++
				mpos++
			}

			continue

		} else if (jpos > k) {
			for (ipos <= j) {
				m[mpos] = data[ipos]
				ipos++
				mpos++
			}
			continue
		}

		//分别从两组的数据头开始，按顺序放入到m中
		if (cf(data[ipos], data[jpos]) < 0) {
			m[mpos] = data[ipos]
			ipos++
			mpos++
		} else {
			m[mpos] = data[jpos]
			jpos++
			mpos++
		}
	}

	copy(data[i:k+1], m[:])
	return 0
}

//i为0 k为size-1
func Mgsort(data []interface{}, i, k int, cf CF) int {
	var j int
	if (i < k) {
		j = (i + k - 1) / 2 //j处于中间位置

		//处理左
		if Mgsort(data, i, j, cf) < 0 {
			return -1
		}

		//处理右
		if Mgsort(data, j+1, k, cf) < 0 {
			return -1
		}

		//合并到一起
		if merge(data, i, j, k, cf) < 0 {
			return -1
		}
	}
	return 0
}
