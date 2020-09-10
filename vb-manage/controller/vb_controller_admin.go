package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	. "vb-server/constant"
	"vb-server/model"
	"vb-server/request"
	"vb-server/response"
	"vb-server/utils/config"
	VBLog "vb-server/utils/log"
	Param "vb-server/utils/param"
	. "vb-server/utils/session"
)

type VBControllerAdmin struct {
	VBControllerBase
	config config.VBConfigServer
}

func (c *VBControllerAdmin) Login(ctx *gin.Context) {

	// request
	loginData := request.VBRAdminLogin{}
	if err:=Param.Bind(ctx,&loginData,"[Admin][绑定参数失败]");err!=nil{
		return;
	}

	// check
	admin := model.NewModelAdmin()
	data, err := admin.DataGetAdmin(loginData.UserName, loginData.PassWord)
	if err != nil {
		VBLog.Print.Info("[Admin][查询用户名密码失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(ADMIN_ERROR_LOGIN_INVALID, map[string]interface{}{}))
		return
	}

	// check fail
	if data.Id == 0 || data.Username == "" {
		//error password
		VBLog.Print.Info("[Admin][登录失败]")
		ctx.JSON(http.StatusOK, response.Error(ADMIN_ERROR_LOGIN_INVALID, map[string]interface{}{}))
		return
	}

	//login success
	//cookie
	VBE, err := ctx.Cookie("VBE")

	//session
	session := VBSession{}
	newId := session.NewId()
	if err == nil {
		newId = VBE
	}

	//set session
	sessionData := VBSessionData{
		SessionKey: newId,
		UserId:     data.Id,
		UserName:   data.Username,
		RealName:   data.Realname,
		LastDate:   time.Now(),
	}

	session.SetSession(newId, sessionData)
	session.UpdateSessionTime(newId, c.config.SessionTime)
	//set cookie
	ctx.SetCookie("VBE", newId, c.config.CookiesTime, "/", c.config.Domain, false, false)

	ctx.JSON(http.StatusOK, response.Success(ADMIN_SUCCESS_LOGIN, map[string]interface{}{}))
}

func (c *VBControllerAdmin) Logout(ctx *gin.Context) {

	//session user
	User := ctx.MustGet("User").(VBSessionData)

	session := VBSession{}
	session.RemoveSession(User.SessionKey)

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"name": User.UserName,
	}))

}


func (c *VBControllerAdmin) Ping(ctx *gin.Context) {


	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"date":time.Now(),
	}))

}

func NewVBControllerAdmin() VBControllerAdmin {

	//config
	serverConfig := config.VBGetConfigServer()

	//new controller
	controllerAdmin := VBControllerAdmin{
		config: serverConfig,
	}
	return controllerAdmin
}
