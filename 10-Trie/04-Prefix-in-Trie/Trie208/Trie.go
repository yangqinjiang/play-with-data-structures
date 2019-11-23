package Trie208

//trie 类
type Trie struct {
	root *node  //根节点不保存字符
	size int
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root:newNode2(),size:0}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
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


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
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


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
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
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

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