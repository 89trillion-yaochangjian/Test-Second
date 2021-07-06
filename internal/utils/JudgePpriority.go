package utils


// 比较运算符栈栈顶top和新运算符newTop的优先级高低

func JudgePpriority(top string, newTop string) bool {
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}