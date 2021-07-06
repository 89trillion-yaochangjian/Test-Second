package router

import (
	"Calculator/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func CalculationCtrl()  {
	r := gin.Default()
	r.GET("/CalculationCtrl", ctrl.CalculationCtrl)
	r.Run(":8080")
}
