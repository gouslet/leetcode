# 206. 反转链表
给你单链表的头节点head，请你反转链表，并返回反转后的链表。

## 示例
### 示例 1：
![](assets/示例1.jpg)
输入：head = [1,2,3,4,5]  
输出：[5,4,3,2,1]  

### 示例 2：
![](assets/示例2.jpg)  
输入：head = [1,2]  
输出：[2,1]   

### 示例 3：

输入：head = []
输出：[]
 

## 提示：

- 链表中节点的数目范围是 [0, 5000]
- -5000 <= Node.val <= 5000
 
## 进阶：
链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

## 解答
### 方法一
**解释**：
找到链表的最后一个有值的节点，即为tail，循环将头结点依次移动至tail的下一个位置，直至tail成为头节点
**代码**
```go
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
```
具体见[reverseList.go](reverseList.go)
### 方法二
**解释**:
递归将头节点移动至尾端  
**代码**
```go
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
```