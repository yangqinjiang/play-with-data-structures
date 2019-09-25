package _4_LinkedList_and_Recursion
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements5(head *ListNode, val int) *ListNode {
	//使用递归
	if nil == head{
		return nil
	}
	//从子集中,删除元素
	head.Next = removeElements5(head.Next,val)
	if head.Val == val{
		return head.Next
	}else{
		return  head
	}
	//return head.Val == val ? head.Next : head
}
func RemoveElements5(head *ListNode, val int) *ListNode {
	return removeElements5(head,val)
}