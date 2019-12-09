package _1_hash_table_basics
//字符串中的第一个唯一字符    387
//https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func FirstUniqChar(s string) int  {
	return firstUniqChar(s)
}
func firstUniqChar(s string) int  {
	freq := [26]int{}
	//累加频率
	for  i:=0;i<len(s);i++ {
		freq[s[i] - 'a'] ++
	}
	//从头开始,判断频繁是一的index
	for  i:=0;i<len(s);i++ {
		if 1 == freq[s[i] - 'a']{
			return i
		}
	}
	return -1
}