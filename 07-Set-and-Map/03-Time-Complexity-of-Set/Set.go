package _3_time_complexity_of_set

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
