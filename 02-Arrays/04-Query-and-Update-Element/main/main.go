package main

import (
	"fmt"
	_4_Query_and_Update_Element "github.com/yangqinjiang/play-with-data-structures/02-Arrays/04-Query-and-Update-Element"
)

func main() {
	arr := _4_Query_and_Update_Element.NewArrayDefault()
	fmt.Println(arr)
	arr.AddFirst(1)
	arr.AddFirst(1)
	arr.Add(2,5)
	arr.AddLast(8)
	arr.Add(4,5)
	arr.AddLast(8)
	fmt.Println(arr.GetSize())
	fmt.Println(arr.Get(2))
	arr.Set(2,10)
	fmt.Println(arr.Get(2) == 10)
	fmt.Println(arr.ToString())
	fmt.Println(arr)
}
