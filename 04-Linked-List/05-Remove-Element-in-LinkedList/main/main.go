package main

import (
	"fmt"
	_5_Remove_Element_in_LinkedList "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/05-Remove-Element-in-LinkedList"
)

func main() {
	ll := _5_Remove_Element_in_LinkedList.NewLinkedList()

	for i:=0;i<5 ;i++  {
		ll.AddFirst(i)
		fmt.Println(ll.ToString())
	}
	ll.Add(2,666)
	if 666 != ll.Get(2){
		panic("should be Get 666")
	}
	if false == ll.Contains(666){
		panic("should be contain 666")
	}
	fmt.Println(ll.ToString())


	ll.Remove(2)
	if true == ll.Contains(666){
		panic("should not be contain 666")
	}
	ll.RemoveFirst()
	ll.RemoveLast()
	fmt.Println(ll.ToString())


}
