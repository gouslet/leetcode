package main

import (
	"fmt"
	"reflect"
)

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	for head != tail {
		temp1 := tail.Next
		temp2 := head
		head = head.Next
		tail.Next = temp2
		temp2.Next = temp1
	}

	return tail
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	left := reverseList2(head.Next)

	tail := left
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = head
	head.Next = nil

	return left
}

func lists2slice(head *ListNode) []int {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
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
		nums1 []int
		res   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
	}
	funcs := []func(head *ListNode) *ListNode{
		reverseList1,
		reverseList2,
	}
	for i, f := range funcs {

		for _, test := range tests {
			l1 := slice2lists(test.nums1)
			res := lists2slice(f(l1))
			fmt.Printf("mergeTwoLists[%d](%v) = %d: ", i, test.nums1, res)
			if !reflect.DeepEqual(test.res, res) {
				fmt.Println("failed")
			} else {
				fmt.Println("passed")
			}
		}
	}
}
