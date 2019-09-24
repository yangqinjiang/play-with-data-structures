package main

import (
	"fmt"
	_4_Query_and_Update_Element "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/04-Query-and-Update-in-LinkedList"
)

func main() {
	ll := _4_Query_and_Update_Element.NewLinkedList()

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


}
