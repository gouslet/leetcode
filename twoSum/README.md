# 167. 两数之和II - 输入有序数组
给定一个已按照非递减顺序排列的整数数组numbers，请你从数组中找出两个数满足相加之和等于目标数target。

函数应该以长度为2的整数数组的形式返回这两个数的下标值。numbers的下标从1开始计数，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length。

你可以假设每个输入只对应唯一的答案 ，而且你不可以重复使用相同的元素。

## 示例 
示例 1：

输入：numbers = [2,7,11,15], target = 9  
输出：[1,2]  
解释：2与7之和等于目标数9。因此 index1 = 1, index2 = 2 。

示例 2：

输入：numbers = [2,3,4], target = 6  
输出：[1,3]

示例 3：

输入：numbers = [-1,0], target = -1  
输出：[1,2]
 

提示：

- 2 <= numbers.length <= 3 * 104
- -1000 <= numbers[i] <= 1000
- numbers 按非递减顺序排列
- -1000 <= target <= 1000
- 仅存在一个有效答案

##解答  
### 方法1

**解释**：从i = 0开始遍历数组，针对每个位置i上的数numbers[i]，以二分搜索在其后所有元素组成的子数组numbers[i+1:]中查找值为target-numbers[i]的位置j，如果找到0 < j < numbers.length，则返回[i+1,j+1]

**代码**：
```go
func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) < 2 || target < numbers[0] {
		return nil
	}

	for i, l := 0, len(numbers); i < l; i++ {
		if j := binsearch(numbers, target-numbers[i], i+1, l-1); j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

func binsearch(nums []int, target, low, high int) int {

	for low <= high {
		mid := low + (high-low)>>1
		if target <= nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low > -1 && low < len(nums) && nums[low] == target {
		return low
	}
	return -1
}
```
具体见[twoSum.go](twoSum.go)