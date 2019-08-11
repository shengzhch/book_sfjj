package common

//希尔排序 多次插入插入排序 充分利用插入排序对处理整体有序的数据的优越性
func insortwithg(data []interface{}, size, g int, cf CF) int {
	var i, j int
	var key interface{}
	for j = g; j < size; j++ {
		key = data[j]

		i = j - g

		//一步一步把key移动到合适的位置，前面的数据是已经排好序的
		for i >= 0 && cf(data[i], key) > 0 {
			data[i+g], data[i] = data[i], data[i+g]
			i -= g
		}
	}
	return 0
}

func ShellSort(data []interface{}, size int, cf CF) {
	var g = make([]int, 0)
	//构造合适的g
	for h := 1; ; {
		if h > size {
			break
		}
		g = append(g, h)
		h = 3*h + 1
	}

	for i := len(g) - 1; i >= 0; i-- {
		insortwithg(data, size, g[i], cf)
	}
}
