package main

import (
	_2_Add_Elements_in_LinkedList "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/02-Add-Elements-in-LinkedList"
)

func main() {
	ll := _2_Add_Elements_in_LinkedList.NewLinkedList()

	for i:=0;i<100 ;i++  {
		ll.AddFirst(i)
	}
}
