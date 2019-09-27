package _9_Non_Recursion_Preorder_Traverse_in_BST

import (
	"container/list"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
//144. 二叉树的前序遍历
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/submissions/
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if nil == root{
		return res
	}
	//使用list模拟Stack
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0{
		cur := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, cur.Val)
		//注意,先压入右子树
		if nil != cur.Right{
			stack.PushBack(cur.Right)
		}
		//再压入右子树
		if nil != cur.Left{
			stack.PushBack(cur.Left)
		}
	}
	return res
}
//供测试使用
func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal(root)
}