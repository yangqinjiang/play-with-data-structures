package _6_Debug_Recursive_Solution

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/

func removeElements(head *ListNode, val int,depth int) *ListNode {
	depthString := generateDepthString(depth)
	fmt.Print(depthString)
	fmt.Println("Call: remove "+strconv.Itoa(val) + " in " + head.ToString())

	//使用递归
	if nil == head{
		fmt.Print(depthString)
		fmt.Println("Return: " + head.ToString())
		return nil
	}
	//从子集中,删除元素
	res := removeElements(head.Next,val,depth  + 1)
	fmt.Print(depthString)
	fmt.Println("After remove " + strconv.Itoa(val) + ": " + head.ToString())
	var ret *ListNode
	if head.Val == val{
		ret = res
	}else{
		head.Next = res
		ret = head
	}
	fmt.Print(depthString)
	fmt.Println("Return " + ret.ToString())
	return ret
}
func RemoveElements(head *ListNode, val int,depth int) *ListNode {
	return removeElements(head,val,depth)
}
//创建字符串
func generateDepthString(depth int) string{
	var sb strings.Builder

	for i := 0;i<depth;i++{
		sb.WriteString("--")
	}
	return sb.String()
}