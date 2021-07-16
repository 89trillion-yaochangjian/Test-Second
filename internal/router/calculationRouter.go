package router

import (
	"Calculator/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func CalculationCtrl() {
	r := gin.Default()
	//计算字符串
	r.POST("/CalculationCtrl", ctrl.CalculationCtrl)
	r.Run(":8080")
}
