package ctrl

import (
	"Calculator/internal/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CalculationCtrl(c *gin.Context) {
	//获取参数
	str := c.Query("str")
	res := handler.CalculationHandler(str)
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}