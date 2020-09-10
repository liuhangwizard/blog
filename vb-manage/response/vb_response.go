package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	. "vb-server/server"
	VBLog "vb-server/utils/log"
)

func Error(message string, data map[string]interface{}) JSON {
	res := JSON{
		"code":    CODE_STATUS_ERROR,
		"message": message,
		"data":    data,
	}
	return res
}

func Success(message string, data map[string]interface{}) JSON {
	res := JSON{
		"code":    CODE_STATUS_SUCCESS,
		"message": message,
		"data":    data,
	}
	return res
}

func Content(code int, message string, data map[string]interface{}) JSON {
	res := JSON{
		"code":    code,
		"message": message,
		"data":    data,
	}
	return res
}

func QueryError(message string, err error, ctx *gin.Context) {
	VBLog.Print.Info(message + err.Error())
	ctx.JSON(http.StatusOK, Error(DATABASE_ERROR_QUERY_FAIL, gin.H{}))
}
