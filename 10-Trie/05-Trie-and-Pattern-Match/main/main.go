package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/10-Trie/05-Trie-and-Pattern-Match"
)

func main() {
	fmt.Println("Trie...")
	trie:= Constructor()
	trie.AddWord("HELLO中方")
	trie.AddWord("HI")
	trie.AddWord("World")
	trie.AddWord("World") //重复的单词
	if 3 != trie.GetSize(){
		panic("单词个数不一致")
	}

	if false == trie.Search("H."){
		panic("应该存在 H. 的单词")
	}
	if false == trie.Search("HELLO.."){
		panic("应该存在 HELLO.. 的单词")
	}
	if false == trie.Search("....d"){
		panic("应该存在 ....d 的单词")
	}
	if true == trie.Search(".d"){
		panic("不应该存在 .d 的单词")
	}
	if false == trie.Search(".....中方"){
		panic("应该存在 .....中方 的单词")
	}
	if false == trie.Search("H....中方"){
		panic("应该存在 H....中方 的单词")
	}
	fmt.Println("OK")
}
