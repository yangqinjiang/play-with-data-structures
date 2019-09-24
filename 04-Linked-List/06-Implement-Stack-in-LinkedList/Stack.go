package _6_implement_stack_in_linkedlist

//stack接口,不支持泛型
type Stack  interface{
	GetSize() int
	IsEmpty() bool
	Push(e int)
	Pop() int
	Peek() int
}
