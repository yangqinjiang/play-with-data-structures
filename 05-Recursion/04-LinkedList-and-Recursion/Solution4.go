package _4_LinkedList_and_Recursion
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements4(head *ListNode, val int) *ListNode {
	//使用递归
	if nil == head{
		return nil
	}
	//从子集中,删除元素
	res := removeElements4(head.Next,val)
	if head.Val == val{
		return  res
	}else{
		head.Next = res
		return  head
	}
}
func RemoveElements4(head *ListNode, val int) *ListNode {
	return removeElements4(head,val)
}