package main

import (
	"fmt"
	_6_loop_queue "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/06-Loop-Queue"
)

func main() {
	//queue := _6_loop_queue.NewArrayQueue(20)
	queue := _6_loop_queue.NewLoopQueue(20)
	for i:=0; i<10;i++  {
		queue.Enqueue(i)
		fmt.Println(queue.ToString())
		if i%3 ==2{
			queue.Dequeue()
			fmt.Println(queue.ToString())
		}
	}


}
