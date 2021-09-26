# 876. 链表的中间结点
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。

## 示例
### 示例 1：

输入：[1,2,3,4,5]  
输出：此列表中的结点 3 (序列化形式：[3,4,5])    
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.  

### 示例 2：

输入：[1,2,3,4,5,6]  
输出：此列表中的结点 4 (序列化形式：[4,5,6])  
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
 

## 提示：

给定链表的结点数介于 1 和 100 之间。

## 解答
### 方法一
**解释**：遍历一次链表，以统计链表长度l，再次遍历以返回第(l + 1)/2个节点
**代码**：
```go
func middleNode(head *ListNode) *ListNode {
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
```
具体见[middleNode.go](middleNode.go)

### 方法二
**解释**：定义两个指针p和q，遍历一次链表，p指针每移动两次，q指针移动一次，同时统计链表长度l，若返回长度为奇数，则返回q.Next，否则返回q
**代码**：
```go
func middleNode(head *ListNode) *ListNode {
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
```
具体见[middleNode.go](middleNode.go)