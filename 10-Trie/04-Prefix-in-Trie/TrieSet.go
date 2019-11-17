package _4_Prefix_in_Trie
/**
实现Set接口的方法
 */
type TrieSet struct {
	trie *Trie
}

func NewTrieSet() *TrieSet  {
	return &TrieSet{trie: NewTrie()}
}
//添加元素到集合内
func (this *TrieSet)Add(e string)  {
	this.trie.Add(e)
}
//查询集合内,是否存在元素e
func (this *TrieSet)Contains(e string) bool{
	return this.trie.Contains(e)
}
//从集合内,删除元素e
func (this *TrieSet)Remove(e string){
	//this.trie.Remove(e)
	panic("未实现此函数")
}
//返回集合的元素个数
func (this *TrieSet)GetSize() int{
	return this.trie.GetSize()
}
//判断集合是否为空
func (this *TrieSet)IsEmpty() bool{
	return 0 == this.trie.GetSize()
}