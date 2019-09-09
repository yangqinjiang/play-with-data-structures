package _3_A_Stack_Problem_in_Leetcode

import (
	"container/list"
)

// leetCode的题目 20
//https://leetcode-cn.com/problems/valid-parentheses/submissions/
func IsValid(s string) bool {
	return isValid(s)
}
func isValid(s string) bool {
	//使用list模拟Stack
	stack := list.New()
	for i:=0;i<len(s) ;i++  {
		c := s[i]

		if c == '(' || c == '[' || c == '{'{
			stack.PushBack(c)//压入
		}else{
			if 0 == stack.Len(){
				return false
			}
			topChar := stack.Remove(stack.Back()).(uint8)//弹出
			//topString := string(topChar)
			//fmt.Println("pf - ",string(c),"-",topString)
			if c == ')' && topChar != '('{
				return false
			}
			if c == ']' && topChar != '['{
				return false
			}
			if c == '}' && topChar != '{'{
				return false
			}
		}
	}
	return stack.Len() == 0
}