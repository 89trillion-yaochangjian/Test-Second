package handler

import (
	"Calculator/internal/service"
	StructInfo "Calculator/internal/structInfo"
)

func CalculationHandler(str string) (StructInfo.Response, error) {
	res, err := service.CalculationService(str)
	return res, err
}
