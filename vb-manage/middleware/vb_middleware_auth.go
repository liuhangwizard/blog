package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	"vb-server/response"
	"vb-server/utils/config"
	. "vb-server/utils/session"
)

var (
	session      *VBSession
	serverConfig config.VBConfigServer
)

func init() {
	session = &VBSession{}
	serverConfig = config.VBGetConfigServer()
}

func VBMiddlewareAuth(context *gin.Context) {

	//get cookie
	VBE, err := context.Request.Cookie("VBE")
	if err != nil {
		//no cookie
		fmt.Println(err)
		context.JSON(http.StatusOK, response.Content(
			CODE_AMDIN_DONOT_LOGIN,
			ADMIN_ERROR_DONOT_LOGIN,
			map[string]interface{}{},
		))
		context.Abort()
		return
	}

	//get session
	user, err := session.GetSession(VBE.Value)
	if user.UserId == 0 || user.UserName == "" {
		//没有这个session
		context.JSON(http.StatusOK, response.Content(
			CODE_ADMIN_TOKEN_INVALID,
			ADMIN_ERROR_SESSION_NONE,
			map[string]interface{}{},
		))
		context.Set("User", user)
		context.Abort()
		return
	}
	//set user data
	context.Set("User", user)
	//session 有效
	//更新时间

	session.UpdateSessionTime(VBE.Value, serverConfig.SessionTime)
	context.SetCookie(
		VBE.Name,
		VBE.Value,
		serverConfig.CookiesTime,
		VBE.Path,
		VBE.Domain,
		VBE.Secure,
		VBE.HttpOnly,
	)

	// request
	context.Next()

}
