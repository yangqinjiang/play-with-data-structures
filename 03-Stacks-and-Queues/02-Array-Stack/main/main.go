package main

import (
	"fmt"
	_2_Array_Stack "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/02-Array-Stack"
)

func main() {
	stack := _2_Array_Stack.NewArrayStackDefault()
	for i:=0;i<5 ;i++  {
		stack.Push(i)
		fmt.Println(stack.ToString())
	}

	stack.Pop()
	fmt.Println(stack.ToString())
}
