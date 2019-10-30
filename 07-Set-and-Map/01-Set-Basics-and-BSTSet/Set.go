package _1_Set_Basics_and_BSTSet

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
