package _6_Debug_Recursive_Solution

import (
    "strconv"
    "strings"
)

type ListNode struct {
     Val int
     Next *ListNode
 }
// 链表节点的构造函数
// 使用arr为参数，创建一个链表，当前的ListNode为链表头结点
func NewListNodeByArray(arr []int) *ListNode {
    if arr == nil || len(arr) == 0{
        panic("arr can not be empty")
    }
    head := &ListNode{Val:arr[0]}
    cur := head
    //循环创建
    for i:=1;i< len(arr);i++{
        cur.Next = &ListNode{Val:arr[i]}
        cur = cur.Next
    }

    return head
}
// 以当前节点为头结点的链表信息字符串
func (this *ListNode)ToString() string  {
    var sb strings.Builder

    cur := this
    for (cur != nil){
        sb.WriteString(strconv.Itoa(cur.Val) + "->")
        cur = cur.Next
    }
    sb.WriteString("NULL")
    return sb.String()
}