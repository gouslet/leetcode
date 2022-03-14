/*
 * File: \mergeTrees\mergeTrees.go                                             *
 * Project: leetcode                                                           *
 * Created At: Tuesday, 2021/09/28 , 23:52:56                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/03/14 , 23:22:21                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	return &TreeNode{root1.Val + root2.Val, mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)}

}
