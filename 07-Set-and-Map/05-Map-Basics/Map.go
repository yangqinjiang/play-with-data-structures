package _5_Map_Basics

//准确的映射接口定义,应该是使用泛型,像java那样 K,V
type Map interface {

	Add(key string ,value int )
	Remove(key string) int //删除key对应的数据,并返回
	Contains(key string) bool
	Get(key string) int //返回key对应的值
	Set(key string,newValue int) //设置新值
	GetSize() int
	IsEmpty() bool
}
