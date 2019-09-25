package _4_LinkedList_and_Recursion
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/

func removeElements2(head *ListNode, val int) *ListNode {

	//删除头部的元素
	for head != nil && head.Val == val{
		head = head.Next  //跳过此元素
	}
	//如果删除头部的全部,返回空
	if nil == head{
		return nil
	}

	//删除中间及后面的元素
	prev := head
	for prev.Next != nil{
		if prev.Next.Val == val{
			prev.Next = prev.Next.Next//跳过此元素
		}else{
			prev = prev.Next
		}
	}
	return head
}
func RemoveElements2(head *ListNode, val int) *ListNode {
	return removeElements2(head,val)
}