package HashTable
import (
	"hash/fnv"
	"strconv"
	// "fmt"
)
const upperTol int = 10
const lowerTol int = 2
const initCapacity = 7 //初始容量
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
	return NewHashTable(initCapacity) //初始容量
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
		//扩容
		if this.size >= upperTol * this.m{
			this.resize(2*this.m)
		}
	}
}
//删除元素
func (this *HashTable)Remove(key string) int{
	var ret int
	_map := this.hashtable[this.hash(key)]

	if _,exist := _map[key];exist{
		ret = _map[key]
		delete(_map,key)
		this.size --
		//缩容
		if this.size < lowerTol * this.m && 
			this.m / 2 >= initCapacity{
			this.resize(this.m / 2)
		}
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
//缩容
func (this *HashTable)resize(newM int){
	//fmt.Printf("%v resize to %v ,%v\n",this.m,newM, newM - this.m)
	newHashTable := make([]map[string]int, newM)
	for i := 0; i < newM; i++ {
		newHashTable[i] = make(map[string]int)
	}
	oldM := this.m
	this.m  = newM
	for i := 0; i < oldM; i++ {
		_map := this.hashtable[i]
		//收集_map的key
		//https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
		keys := make([]string, 0, len(_map))
		for k :=range _map{
			keys = append(keys,	k)
		}
		//再遍历key
		for _,key :=range keys{
			//转移数据到newHashTable
			newHashTable[this.hash(key)][key] = _map[key]
		}
	}
	//更新引用
	this.hashtable = newHashTable
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