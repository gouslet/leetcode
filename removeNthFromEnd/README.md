# 19. 删除链表的倒数第N个结点
给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

## 示例
### 示例 1：
![](assets/19-示例1.jpg)

输入：head = [1,2,3,4,5], n = 2  
输出：[1,2,3,5]  

### 示例 2：

输入：head = [1], n = 1  
输出：[]  

### 示例 3：

输入：head = [1,2], n = 1  
输出：[1]
 
## 提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

## 解答
### 方法一

**解释**：指针p从头节点开始移动n次，然后指针q从头开始，两个指针同步移动，直至指针p指向nil。若在指针p移动n次前已经指向nil，则说明链表节点数量少于n，则报错；若在指针q开始移动前指针p已经指向nil，则说明链表长度为n，则删除头节点；否则删除q节点的下一个节点后返回头节点

**代码**：
```go
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
```
具体见[removeNthFromEnd.go](removeNthFromEnd.go)