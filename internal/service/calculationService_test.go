package service

import "testing"

func TestCalculationService(t *testing.T) {
	CalculationService, err := CalculationService("1+2*3+4*5+6*7")
	t.Log(CalculationService, err)
}
