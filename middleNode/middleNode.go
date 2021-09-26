package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode1(head *ListNode) *ListNode {
	l := 0

	for p := head; p.Next != nil; p = p.Next {
		l++
	}
	q := head
	for i := 0; i < (l+1)/2; i++ {
		q = q.Next
	}

	return q
}

func middleNode2(head *ListNode) *ListNode {
	q := head
	l := 0
	for i, p := 0, head; p.Next != nil; p = p.Next {
		if i != 0 {
			q = q.Next
			i = 0
		} else {
			i++
		}
		l++
	}
	if l%2 == 0 {
		return q
	}
	return q.Next
}

func slice2lists(nums []int) *ListNode {
	l := len(nums)
	if nums == nil || l == 0 {
		return nil
	}
	head := &ListNode{}
	p := head
	for i := 0; i < l-1; i++ {
		p.Val = nums[i]
		p.Next = &ListNode{}
		p = p.Next
	}
	p.Val = nums[l-1]
	p.Next = nil

	return head
}

func main() {
	tests := []struct {
		nums []int
		res  int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
	}
	funcs := []func(head *ListNode) *ListNode{
		middleNode1,
		middleNode2,
	}
	for i, l := 0, len(funcs); i < l; i++ {
		for _, test := range tests {
			head := slice2lists(test.nums)
			val := funcs[i](head).Val
			fmt.Printf("middleNode%d(%v) = %d: ", i+1, test.nums, val)
			if test.res != val {
				fmt.Println("failed")
			} else {
				fmt.Println("passed")
			}
		}
	}
}
