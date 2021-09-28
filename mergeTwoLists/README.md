# 21. 合并两个有序链表
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

## 示例
### 示例 1：
![示例1](assets/示例1.jpg)
输入：l1 = [1,2,4], l2 = [1,3,4]  
输出：[1,1,2,3,4,4]  

### 示例 2：

输入：l1 = [], l2 = []  
输出：[]  

### 示例 3：

输入：l1 = [], l2 = [0]  
输出：[0]  
 

## 提示：

- 两个链表的节点数目范围是 [0, 50]
- -100 <= Node.val <= 100
- l1 和 l2 均按 非递减顺序 排列

## 解答
### 方法一
**解释**：
递归调用函数，每次调用将两个链表的头节点中val值较小者并入新链表
**代码**：
```go
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
```
具体见[mergeTwoLists.go](mergeTwoLists.go)
### 方法二
**解释**：选取两个链表中头节点值较小者为新的头节点，遍历两个链表，依次将较小值的节点连在新的链表的尾部
**代码**：
```go
// 迭代
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
```
具体见[mergeTwoLists.go](mergeTwoLists.go)
