package _3_Recursion_Basics
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/

func removeElements(head *ListNode, val int) *ListNode {

	//删除头部的元素
	for head != nil && head.Val == val{
		delNode := head
		head = head.Next
		delNode.Next = nil
	}
	//如果删除头部的全部,返回空
	if nil == head{
		return nil
	}

	//删除中间及后面的元素
	prev := head
	for prev.Next != nil{
		if prev.Next.Val == val{
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		}else{
			prev = prev.Next
		}
	}
	return head
}
func RemoveElements(head *ListNode, val int) *ListNode {
	return removeElements(head,val)
}