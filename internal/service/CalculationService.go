package service

import (
	StructInfo "Calculator/internal/structInfo"
	"Calculator/internal/utils"
)

func CalculationService(str string) (int, *StructInfo.Response) {
	// 中缀表达式转后缀表达式
	Postfix := utils.InfixToPostfix(str)
	//计算后缀表达式
	res, err := utils.Calculation(Postfix)
	if err != nil {
		return res, StructInfo.PostfixErr
	}
	return res, nil
}
