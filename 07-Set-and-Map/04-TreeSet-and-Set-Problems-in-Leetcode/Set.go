package _4_TreeSet_and_Set_Problems_in_Leetcode

/**
定义集合的接口
 */
type Set interface {
	Add(e string)
	Contains(e string) bool
	Remove(e string)
	GetSize() int
	IsEmpty() bool
}
