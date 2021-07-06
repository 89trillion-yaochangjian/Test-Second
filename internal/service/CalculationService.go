package service

import (
	"Calculator/internal/utils"
	"fmt"
)

func CalculationService(str string)int{
	//判断合法性
	var res int
	isLeg := utils.IsLeg(str)
	if isLeg {
		Postfix := utils.InfixToPostfix(str)
		res = utils.Calculation(Postfix)
	}else {
		fmt.Println("字符串不合法")
	}
	return res
}
