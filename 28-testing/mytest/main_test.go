//cd into the directory with the tests file then run:
//go test
//go test -v (if you want verbose info)
package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4}, 7},
		test{[]int{44, -22}, 22},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("expected", v.answer, "got", x)
		}
	}

}
