package main

import "fmt"

func Hanoi(num int, left, mid, right string) int {
	if num < 1 {
		return 0
	}

	return process(num, left, mid, right, left, right)
}

//递归方式实现汉诺塔问题，规则复杂一点，移到必须经过中间的步骤
func process(num int, left, mid, right, from, to string) int {

	if num == 1 {
		if from == mid || to == mid {
			fmt.Println("move 1 from", from, "to", to)
			return 1
		} else {
			fmt.Println("move 1 from", from, "to", mid)
			fmt.Println("move 1 from", mid, "to", to)
			return 2
		}
	}

	if from == mid || to == mid {
		//从中开始移动 或者移到中
		//三步
		var another string
		if from == left { //from == left;to == mid
			another = right
		} else if to == left { //to == left;from == mid
			another = left
		}

		part1 := process(num-1, left, mid, right, from, another)
		fmt.Println("move", num, "from", from, "to", to)
		part2 := 1
		part3 := process(num-1, left, mid, right, another, to)
		return part1 + part2 + part3
	} else {
		// 从左移到右 或者从右移到左
		//五个步骤

		part1 := process(num-1, left, mid, right, from, to)
		part2 := 1
		fmt.Println("move", num, "from", from, "to", mid)
		part3 := process(num-1, left, mid, right, to, from)
		part4 := 1
		fmt.Println("move", num, "from", mid, "to", to)
		part5 := process(num-1, left, mid, right, from, to)
		return part1 + part2 + part3 + part4 + part5

	}
}



func main() {
	part := Hanoi(2, "left", "mid", "right")
	fmt.Println("part ", part)
}
