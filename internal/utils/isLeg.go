package utils

//判断字符串是否合法

func IsLeg(str string) bool {
	stack := ItemStack{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		switch char {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			continue
		case "-", "+", "*", "/":
			if i == 0 || i == len(str)-1 {
				return false
			}
		//case "(",")":{
		//	if char=="("{
		//		stack.Push(char)
		//	}else if !stack.IsEmpty() {
		//		stack.Pop()
		//	}else {
		//		return false
		//	}
		//}
		default:
			return false
		}
	}
	return stack.IsEmpty()
}
