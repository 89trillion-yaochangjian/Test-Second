package utils

import (
	StructInfo "Calculator/internal/structInfo"
	"strconv"
)

//计算后缀表达式

func Calculation(str string) (int, *StructInfo.Response) {
	var res int
	stack := ItemStack{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		switch char {
		case "+", "-", "*", "/":
			num1, err := strconv.Atoi(stack.Pop())
			if err != nil {
				return res, StructInfo.StrconvErr
			}
			num2, err1 := strconv.Atoi(stack.Pop())
			if err1 != nil {
				return res, StructInfo.StrconvErr
			}
			if char == "+" {
				stack.Push(strconv.Itoa(num1 + num2))
			}
			if char == "-" {
				stack.Push(strconv.Itoa(num1 - num2))
			}
			if char == "*" {
				stack.Push(strconv.Itoa(num1 * num2))
			}
			if char == "/" {
				if num1 == 0 {
					return res, StructInfo.DivisorErr
				}
				stack.Push(strconv.Itoa(num2 / num1))
			}
		default:
			stack.Push(char)
		}
	}
	res, err := strconv.Atoi(stack.Pop())
	if err != nil {
		return res, StructInfo.StrconvErr
	}
	return res, nil
}
