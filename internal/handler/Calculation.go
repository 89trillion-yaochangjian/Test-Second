package handler

import (
	"Calculator/internal/service"
	StructInfo "Calculator/internal/structInfo"
)

func CalculationHandler(str string) (int, *StructInfo.Response) {
	res, err := service.CalculationService(str)
	return res, err
}
