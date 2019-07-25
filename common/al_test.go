package common

import (
	"testing"
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
