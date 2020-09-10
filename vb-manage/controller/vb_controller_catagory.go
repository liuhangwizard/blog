package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	"vb-server/model"
	"vb-server/request"
	"vb-server/response"
	Param "vb-server/utils/param"
)

type VBControllerCatagory struct {
	VBControllerBase
	modelCatagory *model.VBModelCatagory
}


func (c *VBControllerCatagory) GetCatagoryList(ctx *gin.Context){


	//get all
	data,err:=c.modelCatagory.DataGetCatagoryList()
	if err!=nil{
		response.QueryError("[Catagory][获取所有分类失败]",err,ctx)
		return;
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
		"list":data,
	}))

}

func (c *VBControllerCatagory) Delete(ctx *gin.Context){

	queryData:=request.VBRCatagoryDelete{}
	if err:=Param.Bind(ctx,&queryData,"[Catagory][删除分类参数绑定失败]");err!=nil{
		return;
	}

	if err:=c.modelCatagory.Delete(&queryData);err!=nil{
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{}))

}

func (c *VBControllerCatagory) New(ctx *gin.Context){

	queryData:=request.VBRCatagoryNew{}
	if err:=Param.Bind(ctx,&queryData,"[Catagory][新增分类参数绑定失败");err!=nil{
		return;
	}


	if err:=c.modelCatagory.New(&queryData);err!=nil{
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{}))

}


func NewVBControllerCatagory() VBControllerCatagory {

	//new controller
	controllerPermission := VBControllerCatagory{
		modelCatagory:model.NewModelCatagory(),
	}
	return controllerPermission
}
