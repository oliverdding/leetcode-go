package id2476

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=2476 lang=golang
 *
 * [2476] Closest Nodes Queries in a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	res := [][]int{}
	arr := InorderTraversal(root)
	for _, query := range queries {
		min, max := -1, -1
		idx, ok := slices.BinarySearch(arr, query)
		if idx < len(arr) {
			max = arr[idx]
		}
		if !ok { // if arr[idx] == query, arr[idx] is also min
			idx--
		}
		if idx >= 0 {
			min = arr[idx]
		}
		res = append(res, []int{min, max})
	}
	return res
}

func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorderTraversal func(root *TreeNode)
	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		res = append(res, root.Val)
		inorderTraversal(root.Right)
	}
	inorderTraversal(root)
	return res
}

// @lc code=end
