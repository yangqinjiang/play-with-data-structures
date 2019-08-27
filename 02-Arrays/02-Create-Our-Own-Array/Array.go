package _2_Create_Our_Own_Array

type Array struct {
	data []int
	size int
}

//构造器,传入数组的容量capacity构造Array
func NewArray(capacity int) *Array  {
	return &Array{
		data: make([]int,capacity),
		size: 0,
	}
}
//无参数的构造函数,默认数组的容量capacity=10
func NewArrayDefault() *Array  {
	return NewArray(10)
}

//获取数组的容量
func (this *Array)GetCapacity() int  {
	return len(this.data)
}
//获取数组中的元素个数
func (this *Array)GetSize() int {
	return  this.size
}
func (this *Array)IsEmpty() bool  {
	return  this.size == 0;
}