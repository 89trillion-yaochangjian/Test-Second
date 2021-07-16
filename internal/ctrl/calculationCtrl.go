package ctrl

import (
	"Calculator/internal/service"
	result "Calculator/internal/status"
	"Calculator/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//计算字符串
func CalculationCtrl(c *gin.Context) {
	//获取参数
	str := c.PostForm("str")
	//判断字符串是否为空
	if len(str) == 0 {
		c.JSON(http.StatusBadRequest, result.EmptyErr)
		return
	}
	//判断字符串是否合法
	isLeg := utils.IsLeg(str)
	if !isLeg {
		c.JSON(http.StatusBadRequest, result.LegErr)
		return
	}
	//计算字符串
	res, err := service.CalculationService(str)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result.OK.WithData(res))
}
