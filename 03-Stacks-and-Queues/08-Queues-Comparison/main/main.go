package main

import (
	"fmt"
	_8_Queues_Comparison "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/08-Queues-Comparison"
	"math"
	"math/rand"
	"time"
)
// 测试使用q运行opCount个enqueueu和dequeue操作所需要的时间，单位：秒
func testQueue(q _8_Queues_Comparison.Queue,opCount int) float64  {
	
	startTime := time.Now()
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<opCount ;i++  {
		q.Enqueue(r.Intn(math.MaxInt32))
	}
	for i:=0;i<opCount ;i++  {
		q.Dequeue()
	}

	elapsed := time.Since(startTime)
	return elapsed.Seconds()
}

func main() {

	opCount := 100000
	//普通数组
	arrayQueue := _8_Queues_Comparison.NewArrayQueueDefault()
	time1 := testQueue(arrayQueue,opCount)
	fmt.Printf("ArrayQueue, time:  %f s\n",time1)
	//循环队列
	loopQueue := _8_Queues_Comparison.NewLoopQueueDefault()
	time2 := testQueue(loopQueue,opCount)
	fmt.Printf("LoopQueue, time:  %f s\n",time2)




}
