package _5_Array_Queue

//stack接口,不支持泛型
type Queue  interface{
	GetSize() int
	IsEmpty() bool
	Enqueue(e int)
	Dequeue() int
	getFront() int
}
