package _4_Prefix_in_Trie

//trie 的节点
type node struct {
	IsWord bool
	next map[string]*node
}
//两个构造函数
func newNode(isWord bool) *node  {
	return &node{IsWord: isWord,
		next: make(map[string]*node)}
}
func newNode2() *node  {
	return newNode(false)
}

//trie 类
type WordDictionary struct {
	root *node  //根节点不保存字符
	size int
}
/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root:newNode2(),size:0}
}


//获得Trie中存储的单词数量
func (this *WordDictionary)GetSize() int  {
	return this.size
}

//向Trie中添加一个新的单词word
func (this *WordDictionary)AddWord(word string)  {
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

	if !cur.IsWord{
		cur.IsWord = true
		this.size ++
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(word) //
	return this.match(this.root,c2,0)
}
/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) match(node *node,word []rune,index int) bool {
	if index == len(word){
		return node.IsWord
	}
	c2 := word[index]
	c := string(c2) // 从rune转换到 string
	if "." != c{ // 当前c不是 . , 则继续下一个字符
		if nil == node.next[c]{
			return false
		}
		return this.match(node.next[c],word,index + 1)
	}else{ // 当前c是 . ,则遍历其子节点
		keys := this.getKeys1(node.next)
		for _,nextChar := range keys{
			//true 说明子节点找到, false 说明找不到,跳过此循环,找一下子节点
			if this.match(node.next[nextChar],word,index + 1){
				return true
			}
		}
		return false
	}
}
func (this *WordDictionary)getKeys1(m map[string]*node) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}
//单词单词word是否在Trie中
func (this *WordDictionary)Contains(word string) bool  {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(word) //
	cur := this.root
	for i:=0; i< len(c2); i++ {
		c := string(c2[i])
		//单词的某个字符不存在,说明此单词不存在于Trie中
		if nil == cur.next[c] {
			return false
		}
		//跳过下一个节点
		cur = cur.next[c]
	}
	// 找到单词最后一个字符,
	// 并且判断Add的时候,是否作为完整的单词被添加到Trie
	return cur.IsWord
}

//查询是否在Trie中有单词以prefix为前缀
func (this *WordDictionary)IsPrefix(prefix string) bool  {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(prefix) //
	cur := this.root
	for i:=0;i< len(c2);i++  {
		c := string(c2[i])
		//单词的某个字符不存在,说明此单词不存在于Trie中
		if nil == cur.next[c] {
			return false
		}
		//跳过下一个节点
		cur = cur.next[c]
	}
	//如果运行到这, 说明prefix存在于trie中
	return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */