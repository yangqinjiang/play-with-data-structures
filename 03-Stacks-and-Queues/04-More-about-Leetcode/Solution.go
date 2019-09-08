package _4_MORE_ABOUT_LEETCODE

// leetCode的题目,使用自定义的stack
func IsValid(s string) bool {
	return isValid(s)
}
func isValid(s string) bool {
	if 0 == len(s){
		return false
	}
	stack :=  NewArrayStack(20)
	for i:=0;i< len(s); i++ {
		c := s[i]
		//左括号,压入堆栈
		if c == '(' || c == '[' || c=='{'{
			stack.Push(int(c))
		}else{
			//为空,则返回false
			if stack.IsEmpty(){
				return false
			}
			
			topChar := stack.Pop()
			if c == ')' && topChar != int('(') {
				return false
			}
			if c == ']' && topChar != int('[') {
				return false
			}
			if c == '}' && topChar != int('{') {
				return false
			}
		}

	}
	return stack.IsEmpty()
}