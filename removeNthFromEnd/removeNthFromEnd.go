package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return nil
	}

	p, q := head, head

	for i := 0; i < n; i++ {
		if p == nil {
			panic(fmt.Sprintf("removeNthFromEnd: list doesn't have %d nodes\n", n))
		}
		p = p.Next
	}
	if p == nil {
		return head.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	temp := q.Next
	if temp == nil {
		return nil
	}
	q.Next = temp.Next
	temp.Next = nil

	return head
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

func lists2slice(head *ListNode) []int {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func main() {
	tests := []struct {
		nums []int
		n    int
		res  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{1, 2, 4, 5, 6}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, test := range tests {
		head := slice2lists(test.nums)
		head = removeNthFromEnd(head, test.n)
		fmt.Printf("removeNthFromEnd(%v,%d) = %v: ", test.nums, test.n, lists2slice(head))
		if !reflect.DeepEqual(slice2lists(test.res), head) {
			fmt.Println("failed")
		} else {
			fmt.Println("passed")
		}
	}
}
