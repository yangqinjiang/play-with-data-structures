package main

import (
	"fmt"
	_7_Implementation_of_Loop_Queue "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/07-Implementation-of-Loop-Queue"
)

func main() {

	queue := _7_Implementation_of_Loop_Queue.NewLoopQueue(10)
	for i:=0; i<30;i++  {
		queue.Enqueue(i)
		fmt.Println(queue.ToString())
		if i%3 ==2{
			//fmt.Println("Dequeueing ",i)
			queue.Dequeue()
			fmt.Println(queue.ToString())
		}
	}


}
