package ctrl

import (
	"Calculator/internal/handler"
	result "Calculator/internal/structInfo"
	"Calculator/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CalculationCtrl(c *gin.Context) {
	//获取参数
	str := c.Query("str")
	//判断字符串是否为空
	if len(str) == 0 {
		c.JSON(http.StatusInternalServerError, result.EmptyErr)
		return
	}
	//判断字符串是否合法
	isLeg := utils.IsLeg(str)
	if !isLeg {
		c.JSON(http.StatusInternalServerError, result.LegErr)
		return
	}
	//计算字符串
	res, err := handler.CalculationHandler(str)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, result.DivisorErr)
		return
	}
	c.JSON(http.StatusOK, result.OK.WithData(res.Data))
}
