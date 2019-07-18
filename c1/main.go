package main

import (
	"fmt"
	"github.com/shengzhch/book_sfjj/common"
	)

func Jiecheng(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Jiecheng(n-1)
}

func Jiecheng_weidigui(n int, a int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return Jiecheng_weidigui(n-1, n*a)
}

type A struct {
	aa string
}

func main() {
	rel := common.Hashpjw("s",100)
	fmt.Println("result ",rel)
	return


	var aa *A

	if aa == nil {
		fmt.Println("Done")
		aa = &A{
			"string",
		}
	}

	fmt.Println("aa.aa ", aa.aa)

	return
	fmt.Println(Jiecheng(4))
	fmt.Println(Jiecheng_weidigui(4, 1))
}
