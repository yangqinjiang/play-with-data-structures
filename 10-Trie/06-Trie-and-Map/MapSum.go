package _6_Trie_and_Map

///https://leetcode-cn.com/problems/map-sum-pairs/
//677. 键值映射
type MapSum struct {
	root *node  //根节点不保存字符
	size int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{root:newNode2(),size:0}
}


func (this *MapSum) Insert(word string, val int)  {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(word) //
	//fmt.Println("add word len =",len(c2))
	cur := this.root
	for i:=0;i< len(c2);i++  {
		c := string(c2[i])
		//fmt.Println(c,c2[i])
		//无此节点,补充一下
		if nil == cur.next[c] {
			cur.next[c] = newNode2()
		}
		//跳过下一个节点
		cur = cur.next[c]
	}

	cur.value = val
}


func (this *MapSum) Sum(prefix string) int {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(prefix) //
	cur := this.root
	for i:=0; i< len(c2); i++ {
		c := string(c2[i])
		//单词的某个字符不存在,说明此单词不存在于Trie中
		if nil == cur.next[c] {
			return 0
		}
		//跳过下一个节点
		cur = cur.next[c]
	}

	//递归
	return this._sum(cur)
}
func (this *MapSum)_sum(node *node) int {
	res := node.value
	keys := this.getKeys1(node.next)
	//for 循环,包含了递归的终止条件
	for _,nextChar := range keys{
		res += this._sum(node.next[nextChar])
	}
	return res;
}
func (this *MapSum)getKeys1(m map[string]*node) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
//trie 的节点
type node struct {
	value int   //保存其值
	next map[string]*node
}
//两个构造函数
func newNode(value int) *node  {
	return &node{value: value,
		next: make(map[string]*node)}
}
func newNode2() *node  {
	return newNode(0)
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
