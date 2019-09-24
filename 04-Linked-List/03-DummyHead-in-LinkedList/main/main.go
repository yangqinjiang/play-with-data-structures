package main

import (
	_3_DummyHead_in_LinkedList "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/03-DummyHead-in-LinkedList"
)

func main() {
	ll := _3_DummyHead_in_LinkedList.NewLinkedList()

	for i:=0;i<100 ;i++  {
		ll.AddFirst(i)
	}
}
