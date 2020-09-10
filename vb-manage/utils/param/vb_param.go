package Param

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vb-server/response"
	VBLog "vb-server/utils/log"
)

func Bind(ctx *gin.Context,data interface{},errorMessage string) error {

	//外面用了取地址符号 这里不需要了 不然binding required 无效(即无论参数如何都绑定成功 err永远nil)
	err := ctx.ShouldBind(data)
	if err != nil {
		VBLog.Print.Info(errorMessage + err.Error())
		ctx.JSON(http.StatusOK, response.Error("参数错误", map[string]interface{}{}))
		ctx.Abort()
		return err
	}
	return nil
}