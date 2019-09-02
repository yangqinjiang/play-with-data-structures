package _2_Array_Stack

//stack接口,不支持泛型
type Stack  interface{
	GetSize() int
	IsEmpty() bool
	Push(e int)
	Pop() int
	Peek() int
}
