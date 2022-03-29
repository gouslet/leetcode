# 剑指 Offer II 074. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

# 示例：
## 示例 1：

### 输入：
intervals = [[1,3],[2,6],[8,10],[15,18]]
### 输出：
[[1,6],[8,10],[15,18]]
### 解释：
区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

## 示例 2：

### 输入：
intervals = [[1,4],[4,5]]
### 输出：
[[1,5]]
### 解释：
区间 [1,4] 和 [4,5] 可被视为重叠区间。

# 解答
## 解释
- 首先将区间按照左端点位置排序，然后用区域右端点的位置关系去判断合并或是丢弃
- 取第一个区间为t，依次用后续的区间右端点的位置与之进行比较
- 如果当前区间的左端点在t区间的右端点之后，那么它们不会重合，可以直接将t区间加入数组 res 的末尾，并令t=当前区间
- 否则，它们重合，需要用当前区间的右端点更新t区间的右端点，将其置为二者的较大值。
## 代码
```go
func merge(intervals [][]int) [][]int {
	if intervals == nil {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t := intervals[0]
	res := [][]int{}

	for i, l := 1, len(intervals); i < l; i++ {
		if inter := intervals[i]; inter[0] > t[1] {
			res = append(res, t)
			t = inter
		} else if inter[1] > t[1] {
			t[1] = inter[1]
		}
	}
	res = append(res, t)

	return res
}
```
## 复杂度分析
### 时间复杂度：
O(nlogn)，其中 n 为区间的数量。除去排序的开销，我们只需要一次线性扫描，所以主要的时间开销是排序的 O(nlogn)。

### 空间复杂度：
O(logn)，其中 n 为区间的数量。这里计算的是存储答案之外，使用的额外空间。O(logn) 即为排序所需要的空间复杂度。
