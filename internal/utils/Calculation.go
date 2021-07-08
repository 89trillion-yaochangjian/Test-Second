package utils

import (
	"errors"
	"strconv"
)

//计算后缀表达式

func Calculation(str string) (int, error) {
	var res int
	stack := ItemStack{}
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		switch char {
		case "+", "-", "*", "/":
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
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
					err := errors.New("除数不能为零")
					return res, err
				}
				stack.Push(strconv.Itoa(num2 / num1))
			}
		default:
			stack.Push(char)
		}
	}
	res, err := strconv.Atoi(stack.Pop())
	return res, err
}
