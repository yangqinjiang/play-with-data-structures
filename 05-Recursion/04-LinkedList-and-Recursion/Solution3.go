package _4_LinkedList_and_Recursion
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements3(head *ListNode, val int) *ListNode {

	//使用虚拟头结点
	dummyHead := &ListNode{Next:head,Val:-1}


	//删除虚拟头结点后面的元素
	prev := dummyHead
	for prev.Next != nil{
		if prev.Next.Val == val{
			prev.Next = prev.Next.Next//跳过此元素
		}else{
			prev = prev.Next
		}
	}
	return dummyHead.Next
}
func RemoveElements3(head *ListNode, val int) *ListNode {
	return removeElements3(head,val)
}