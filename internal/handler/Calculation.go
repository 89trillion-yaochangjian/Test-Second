package handler

import (
	"Calculator/internal/service"
)

func CalculationHandler(str string)int{
	return service.CalculationService(str)
}