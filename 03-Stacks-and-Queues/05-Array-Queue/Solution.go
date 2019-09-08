package _5_Array_Queue

import "container/list"

/// Leetcode 102. Binary Tree Level Order Traversal
/// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/// 二叉树的层序遍历
///
/// 二叉树的层序遍历是一个典型的可以借助队列解决的问题。
/// 该代码主要用于使用Leetcode上的问题测试我们的ArrayQueue。
/// 对于二叉树的层序遍历，这个课程后续会讲到。
/// 届时，同学们也可以再回头看这个代码。
/// 不过到时，大家应该已经学会自己编写二叉树的层序遍历啦：）
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//双向队列
	queue := list.New()
	//头部插入
	queue.PushFront(root)
	//进行广度搜索
	for queue.Len() > 0  {
		var currentLevel []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++  {
			//尾部移除,转换成 TreeNode类型的对象
			node := queue.Remove(queue.Back()).(*TreeNode)
			//同一层次,增加node.val
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		res = append(res, currentLevel)
	}
	return res
	//
	//作者：xsplus
	//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/shen-du-you-xian-yan-du-you-xian-go-by-xsplus/
	//来源：力扣（LeetCode）
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


	//res := [][]int{}
	//if root == nil{
	//	return res
	//}
	//// 我们使用list来做为我们的先入先出的队列
	////queue := NewArrayQueueDefault()
	////queue.Enqueue()
	////queue := list.New()
	////queue.PushFront()
	//
	//return res
}
