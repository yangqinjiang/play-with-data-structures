package HashTable
import (
	"hash/fnv"
	"strconv"
	// "fmt"
)
//哈希表
type HashTable struct{
	hashtable []map[string]int
	size,m int
}

func NewHashTable(M int) *HashTable  {
	ht := make([]map[string]int, M)
	for i := 0; i < M; i++ {
		ht[i] = make(map[string]int)
	}
	return &HashTable{m:M,size:0,hashtable:ht}
}
func NewHashTable2() *HashTable{
	return NewHashTable(97) //素数
}

//计算哈希值
func (this *HashTable)hash(key string) int {

	return int(((HashCode(key) & 0x7fffffff) % uint64(this.m)))
}
func (this *HashTable)GetSize() int{
	return this.size
}
//添加元素
func (this *HashTable)Add(key string,value int) {
	_map := this.hashtable[this.hash(key)]
	if _,exist := _map[key];exist{
		_map[key] = value
	}else{
		_map[key] = value
		this.size ++
	}
}
//删除元素
func (this *HashTable)Remove(key string) int{
	var ret int
	_map := this.hashtable[this.hash(key)]

	if _,exist := _map[key];exist{
		delete(_map,key)
		this.size --
	}
	return ret
}
//更新元素
func (this *HashTable)Set(key string,value int)  {
	
	_map := this.hashtable[this.hash(key)]

	if _,exist := _map[key];!exist{
		panic(key + "doesn't exist!")
	}
	_map[key] = value
}
//是否包含元素
func (this *HashTable)Contains(key string) bool{
	
	_,exist := this.hashtable[this.hash(key)][key]
	return exist

}
//获取元素
func (this *HashTable)Get(key string) int{
	return this.hashtable[this.hash(key)][key]
}

//计算哈希值
func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64),'f',6,64)
	}
	h := fnv.New64a()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}