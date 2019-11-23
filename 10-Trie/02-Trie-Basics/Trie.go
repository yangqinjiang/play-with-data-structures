package _2_Trie_Basics

import "fmt"

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
type Trie struct {
	root *node  //根节点不保存字符
	size int
}

func NewTrie() *Trie  {
	return &Trie{root:newNode2(),size:0}
}

//获得Trie中存储的单词数量
func (this *Trie)GetSize() int  {
	return this.size
}

//向Trie中添加一个新的单词word
func (this *Trie)Add(word string)  {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	c2:=[]rune(word) //
	fmt.Println("add word len =",len(c2))
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