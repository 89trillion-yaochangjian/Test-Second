package utils

import (
	"strconv"
)

//计算后缀表达式

func Calculation(str string) int{
	stack := ItemStack{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		switch char {
		case "+","-","*","/":
			num1,_ := strconv.Atoi(stack.Pop())
			num2,_ := strconv.Atoi(stack.Pop())
			if char == "+" {
				stack.Push(strconv.Itoa(num1+num2))
			}
			if char == "-" {
				stack.Push(strconv.Itoa(num1-num2))
			}
			if char == "*" {
				stack.Push(strconv.Itoa(num1*num2))
			}
			if char == "/" {
				stack.Push(strconv.Itoa(num1/num2))
			}
		default:
			stack.Push(char)
		}
	}
	res,_:=strconv.Atoi(stack.Pop())
	return res
}