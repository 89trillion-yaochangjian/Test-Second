package utils

// 中缀表达式转后缀表达式

func InfixToPostfix(fix string) string{
	stack := ItemStack{}
	len := len(fix)
	postfix := ""
	//遍历字表达式
	for i := 0; i < len; i++ {
		char := string(fix[i])
		switch char {
		case "(":
			stack.Push("(")
		case ")":
			for !stack.IsEmpty(){
				prochar := stack.Top()
				if prochar == "(" {
					stack.Pop()
					break
				}
				postfix += prochar
				stack.Pop()
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":{
			postfix += char
		}
		default:
			for !stack.IsEmpty() {
				top := stack.Top()
				if JudgePpriority(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}