# 733. 图像渲染
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在0到65535之间。

给你一个坐标(sr, sc)表示图像渲染开始的像素值（行，列）和一个新的颜色值newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

## 示例
### 示例 1:

输入:   
- image = [[1,1,1],[1,1,0],[1,0,1]]  
- sr = 1, sc = 1, newColor = 2  

输出: [[2,2,2],[2,2,0],[2,0,1]]  
解析: 
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。  
注意，右下角的像素没有更改为2，因为它不是在上下左右四个方向上与初始点相连的像素点。

## 注意:

- image 和 image[0] 的长度在范围 [1, 50] 内。
- 给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
- image[i][j] 和 newColor 表示的颜色值在范围[0, 65535]内。

## 解答
### 方法一
**解释**：  
以深度优先搜索的方式查找满足条件的像素点，修改其像素值

**代码**：
```go
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	l := len(image)
	if l == 0 {
		return image
	}
	c := len(image[0])
	if c == 0 {
		return image
	}

	color := image[sr][sc]
	flag := [][]bool{}
	for i := 0; i < l; i++ {
		flag = append(flag, make([]bool, c))
	}
	image = dfs(image, sr, sc, flag, color, newColor)

	return image
}

func dfs(image [][]int, sr int, sc int, flag [][]bool, color, newColor int) [][]int {
	for _, a := range adj(image, sr, sc) {
		if len(a) == 0 {
			return image
		}
		if !flag[a[0]][a[1]] && image[a[0]][a[1]] == color {
			image[a[0]][a[1]] = newColor
			flag[a[0]][a[1]] = true
			image = dfs(image, a[0], a[1], flag, color, newColor)
		}
	}
	flag[sr][sc] = true
	image[sr][sc] = newColor

	return image
}

func adj(image [][]int, sr, sc int) [][]int {
	res := [][]int{}
	lmax := len(image) - 1
	if lmax < 0 {
		return res
	}
	cmax := len(image[0]) - 1
	if cmax < 0 {
		return res
	}
	if lmax == 0 {
		if cmax == 0 {
			return res
		} else {
			if sc == 0 {
				return [][]int{{0, 1}}
			} else if sc == cmax {
				return [][]int{{0, cmax - 1}}
			} else {
				return [][]int{{0, sc - 1}, {0, sc + 1}}
			}
		}
	}

	if cmax == 0 {
		if lmax == 0 {
			return res
		} else {
			if sr == 0 {
				return [][]int{{1, 0}}
			} else if sr == lmax {
				return [][]int{{lmax - 1, 0}}
			} else {
				return [][]int{{sr - 1, 0}, {sr + 1, 0}}
			}
		}
	}

	if sr == 0 {
		if sc == 0 {
			res = append(res, []int{0, 1}, []int{1, 0})
		} else if sc == cmax {
			res = append(res, []int{0, cmax - 1}, []int{1, cmax})
		} else {
			res = append(res, []int{0, sc - 1}, []int{0, sc + 1}, []int{1, sc})
		}
	} else if sr == lmax {
		if sc == 0 {
			res = append(res, []int{lmax, 1}, []int{lmax - 1, 0})
		} else if sc == cmax {
			res = append(res, []int{lmax, cmax - 1}, []int{lmax - 1, cmax})
		} else {
			res = append(res, []int{lmax, sc - 1}, []int{lmax, sc + 1}, []int{lmax - 1, sc})
		}
	} else {
		if sc == 0 {
			res = append(res, []int{sr, 1}, []int{sr - 1, 0}, []int{sr + 1, 0})
		} else if sc == cmax {
			res = append(res, []int{sr, cmax - 1}, []int{sr - 1, cmax}, []int{sr + 1, cmax})
		} else {
			res = append(res, []int{sr, sc - 1}, []int{sr, sc + 1}, []int{sr - 1, sc}, []int{sr + 1, sc})
		}
	}

	return res
}
```