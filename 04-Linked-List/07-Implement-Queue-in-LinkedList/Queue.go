package _7_Implement_Queue_in_LinkedList

//stack接口,不支持泛型
type Queue  interface{
	GetSize() int
	IsEmpty() bool
	Enqueue(e int)
	Dequeue() int
	GetFront() int
}
