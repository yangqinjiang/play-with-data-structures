package _6_Priority_Queue

//实现Queue接口
type PriorityQueue struct {
	maxHeap *MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{maxHeap:NewMaxHeap()}
}
func (this *PriorityQueue)GetSize() int  {
	return this.maxHeap.Size()
}
func (this *PriorityQueue)IsEmpty() bool  {
	return this.maxHeap.IsEmpty()
}

func (this *PriorityQueue)getFront(e int) int {
	return this.maxHeap.FindMax()
}
func (this *PriorityQueue)Enqueue(e int) {
	this.maxHeap.Add(e)
}
func (this *PriorityQueue)Dequeue() int {
	return  this.maxHeap.ExtractMax()
}