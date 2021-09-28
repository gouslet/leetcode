package main

import (
	"fmt"
	"reflect"
)

/**
 * Definition for singly-linked list.
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		} else {
			return l2
		}
	} else {
		if l2 == nil {
			return l1
		}
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		} else {
			return l2
		}
	} else {
		if l2 == nil {
			return l1
		}
	}
	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = head.Next
		head.Next = nil
	} else {
		head = l2
		l2 = head.Next
		head.Next = nil
	}
	cur := head
	for l1 != nil {
		if l2 == nil {
			cur.Next = l1
			break
		}
		if l1.Val > l2.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
			cur.Next = nil
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
			cur.Next = nil
		}
	}

	if l2 != nil {
		cur.Next = l2
	}

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
		nums2 []int
		res   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 7, 8}, []int{3, 4, 5, 6, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	funcs := []func(s1 *ListNode, s2 *ListNode) *ListNode{
		mergeTwoLists1,
		mergeTwoLists2,
	}
	for i, f := range funcs {

		for _, test := range tests {
			l1 := slice2lists(test.nums1)
			l2 := slice2lists(test.nums2)
			res := lists2slice(f(l1, l2))
			fmt.Printf("mergeTwoLists[%d](%v,%v) = %d: ", i, test.nums1, test.nums2, res)
			if !reflect.DeepEqual(test.res, res) {
				fmt.Println("failed")
			} else {
				fmt.Println("passed")
			}
		}
	}
}
