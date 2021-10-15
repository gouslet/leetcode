package main

import (
	"fmt"
	"testing"
)

func TestPopulating_Next_Right_Pointers_in_Each_Node(t *testing.T) {
	tests := []struct {
		eles []int
		res  []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{`1`, `#`}},
		{[]int{1, 2, 3}, []string{`1`, `#`, `2`, `3`, `#`}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []string{`1`, `#`, `2`, `3`, `#`, `4`, `5`, `6`, `7`, `#`}},
	}
	for _, test := range tests {
		tree := NewTreeFrom(test.eles)
		tree = populating_Next_Right_Pointers_in_Each_Node(tree)
		fmt.Printf("populating_Next_Right_Pointers_in_Each_Node(%v) = %v: ", test.eles, tree)
		got := fmt.Sprintf("%v", tree)
		want := fmt.Sprintf("%v", test.res)
		if got != want {
			t.Errorf("failed: got %v,want %v\n", tree, test.res)
		}
	}
}
