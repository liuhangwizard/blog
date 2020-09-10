package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	"vb-server/model"
	"vb-server/response"
	. "vb-server/utils/session"
)

type VBControllerInfo struct {
	VBControllerBase
	modelTag *model.VBModelTag
	modelCatagory *model.VBModelCatagory
	modelPermission *model.VBModelPermission
}

func (c *VBControllerInfo) GetUserInfo(ctx *gin.Context) {


	//session user
	userData := ctx.MustGet("User").(VBSessionData)



	permission:=c.modelPermission
	permissionData,err:=permission.DataGetUserPermission(userData.UserId)
	if err!=nil{
		response.QueryError("[Info][查询用户权限失败]]",err,ctx)
		return;
	}

	catagory:=c.modelCatagory
	catagoryData,err:=catagory.DataGetCatagoryList()
	if err!=nil{
		response.QueryError("[Info][查询标签失败]]",err,ctx)
		return;
	}


	tag:=c.modelTag
	tagData,err:=tag.DataGetTagList()
	if err!=nil{
		response.QueryError("[Info][查询标签失败]]",err,ctx)
		return;
	}
	

	ctx.JSON(http.StatusOK,response.Success(COMMON_SUCCESS,gin.H{
		"name":userData.RealName,
		"config":permissionData.Config,
		"catagoryList":catagoryData,
		"tagList":tagData,
	}))
}


func NewVBControllerInfo() VBControllerInfo {

	//new controller
	controllerInfo := VBControllerInfo{
		modelTag:model.NewModelTag(),
		modelCatagory:model.NewModelCatagory(),
		modelPermission:model.NewModelPermission(),
	}
	return controllerInfo
}
