package common

import "testing"

var cf = func(key1, key2 interface{}) bool {
	if key1.(int) >= key2.(int) {
		return true
	} else {
		return false
	}
}

func TestInsort(t *testing.T) {
	a := []interface{}{
		1, 5, 7, 8, 3, 2, 4, 6,
	}

	Insort(a, 8, cf)

	t.Log("rel ", a)
}
