package common

import (
	"testing"
	"fmt"
)

var cf = func(key1, key2 interface{}) int {
	if key1.(int) > key2.(int) {
		return 1
	} else if key1.(int) < key2.(int) {
		return -1
	} else {
		return 0
	}
}

func TestInsort(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6,
	}

	Insort(a, 8, cf)

	t.Log("rel ", a)
}

func TestQksort(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45,
	}
	l := len(a) - 1

	Qksort1(a, 0, l, cf)

	t.Log("rel ", a)
}

func TestQksort2(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45,
	}
	l := len(a) - 1

	Qksort2(a, 0, l, cf)

	t.Log("rel ", a)
}

func TestShellSort(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 36, 25,
	}
	l := len(a)

	ShellSort(a, l, cf)

	t.Log("rel ", a)
}

func TestMgsort(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 36, 25,
	}
	l := len(a) - 1

	Mgsort(a, 0, l, cf)

	t.Log("rel ", a)
}

func TestCtsort(t *testing.T) {
	a := []int{
		1, 5, 7, -8, 3, 2, 4, 6, 8, 5, 0, 4, 11, 45, 36, 25,
	}
	l := len(a)

	Ctsort(a, l, -8, 45)

	t.Log("rel ", a)
}

func TestRxsort(t *testing.T) {
	a := []int{
		302, 253, 11, 901, 529, 102,
	}
	l := len(a)

	Rxsort(a, l, 3, 10)

	t.Log("rel ", a)
}

func TestBisearch(t *testing.T) {
	a := []interface{}{
		0, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 8, 11, 45,
	}
	l := len(a)

	rel := Bisearch(a, 5, l, cf)

	t.Log("rel ", rel)
}

func TestBiSearchTree(t *testing.T) {
	var bst = &BiSearchTree{}
	bst.Init(Compare_int)
	bst.Add(15)
	bst.Add(6)
	bst.Add(18)
	bst.Add(3)
	bst.Add(7)
	bst.Add(17)
	bst.Add(20)
	bst.Add(2)
	bst.Add(4)
	bst.Add(13)
	bst.Add(9)
	bst.Add(14)

	fmt.Printf("The nodes of the BiSearchTree is: ")
	bst.InOrderTravel()
	fmt.Printf("\n")

	fmt.Printf("The deepth of the tree is: %d\n", bst.GetDeepth())
	fmt.Printf("The min is: %v\n", bst.GetMin())
	fmt.Printf("The max is: %v\n", bst.GetMax())

	if bst.Search(17) != nil {
		fmt.Printf("The 17 exists.\n")
	}

	fmt.Printf("root node's predecessor is: %v\n", bst.GetPredecessor(bst.GetRoot()).data)
	fmt.Printf("root node's successor is: %v\n", bst.GetSuccessor(bst.GetRoot()).data)

	bst.Delete(13)
	fmt.Printf("Nodes after delete the 13: ")
	bst.InOrderTravel()
	fmt.Printf("\n")
}
