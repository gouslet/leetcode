# 704. 二分查找
给定一个n个元素有序的（升序）整型数组nums和一个目标值target，写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回-1。

## 示例
### 示例1:

输入:nums = [-1,0,3,5,9,12], target = 9 
输出:4  
解释:9出现在nums中并且下标为4  

### 示例2:

输入:nums = [-1,0,3,5,9,12], target = 2  
输出: -1  
解释:2不存在nums中因此返回-1  

## 提示：
- 你可以假设nums中的所有元素是不重复的。
- n将在[1, 10000]之间。
- nums的每个元素都将在[-9999, 9999]之间。

## 解答
### 方法一（递归）
**解释**：
- 辅助函数原型为`func binsearch(nums []int,target,low,high int) int`
- 取mid为low和high的中间
  - 若nums[mid]大于target，则返回`binsearch(nums,target,low,mid - 1)`
  - 若nums[mid]小于target，则返回`binsearch(nums,target,mid + 1,high)`  
  - 否则返回mid
- 函数原型为`func search(nums []int,target int) int`可实现为调用`binsearch(nums,target,0,len(nums) - 1)`

**代码**： 
```go
func binsearch(nums []int, target,low,high int) int {
    if nums == nil || nums[low] > target || nums[high] < target {
        return -1
    }
    mid := low + (high - low) >> 1

    if target > nums[mid] {
        return binsearch(nums,target,mid + 1,high)
    } else if target < nums[mid] {
        return binsearch(nums,target,low,mid - 1)
    } else {
        return mid
    }
}
```

### 方法二（迭代）
**解释**：
- 函数原型为`func search2(nums []int,target int) int`
- 取low和high初始化为0和`len(nums) - 1`
- 当low不大于high的时候，循环
  - 取mid为low和high的中间
    - 若nums[mid]大于target，则`high = mid - 1`
    - 若nums[mid]小于target，则`low = mid + 1`  
    - 否则返回mid
- 循环正常结束则返回-1

**代码**： 
```go
func search2(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	if l == 0 {
		return -1
	}

	for low <= high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
```
具体见[binsearch.go](binsearch.go)

### 方法三（迭代）
**解释**：
- 函数原型为`func search3(nums []int,target int) int`
- 取low和high初始化为0和`len(nums) - 1`
- 当low小于high的时候，循环
  - 取mid为low和high的中间
    - 若nums[mid]小于target，则`high = mid + 1`
    - 否则`high = mid`
- 循环结束后，如果`target == nums[low]`，则返回low，否则返回-1

**代码**： 
```go
func search3(nums []int, target int) int {
	l := len(nums)
	low, high := 0, l-1
	if l == 0 {
		return -1
	}

	for low < high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if target == nums[low] {
		return low
	} else {
		return -1
	}
}
```
具体见[binsearch.go](binsearch.go)

## 优化
因为
- n将在[1, 10000]之间。
- nums的每个元素都将在[-9999, 9999]之间。

故可以使用函数原型`func search(nums []int16,target int16) int`