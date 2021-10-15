package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (root *Node) String() string {
	strs := []string{}
	for p := root; p != nil; p = p.Left {
		strs = append(strs, fmt.Sprintf("%d", p.Val))
		n := p.Next
		for n != nil {
			strs = append(strs, fmt.Sprintf("%d", n.Val))
			n = n.Next
		}
		if n == nil {
			strs = append(strs, `#`)
		}
	}
	return fmt.Sprintf("%v", strs)
}

func NewTreeFrom(eles []int) *Node {
	if l := len(eles); l&(l+1) != 0 {
		panic(fmt.Sprintf("error: can't make up a Full Binary Tree with %v\n", eles))
	}
	return ll(eles)
}

func ll(eles []int) *Node {
	l := len(eles)
	nodes := []*Node{}
	if l == 0 {
		return nil
	} else {
		root := &Node{eles[0], nil, nil, nil}
		if l == 1 {
			return root
		}

		nodes = append(nodes, root)

		eles = eles[1:]
		for len(nodes) != 0 && len(eles) != 0 {
			nodes[0].Left = &Node{eles[0], nil, nil, nil}
			nodes[0].Right = &Node{eles[1], nil, nil, nil}
			nodes = append(nodes, nodes[0].Left, nodes[0].Right)
			nodes = nodes[1:]
			eles = eles[2:]
		}
		return root
	}
}

func populating_Next_Right_Pointers_in_Each_Node(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	l, r := root.Left, root.Right
	populating_Next_Right_Pointers_in_Each_Node(l)
	populating_Next_Right_Pointers_in_Each_Node(r)

	for l != nil {
		l.Next = r
		l = l.Right
		r = r.Left
	}

	return root
}

func main() {
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
		if !reflect.DeepEqual(fmt.Sprintf("%v", tree), fmt.Sprintf("%v", test.res)) {
			fmt.Printf("failed: want %v\n", test.res)
		} else {
			fmt.Println("passed")
		}
	}
}
