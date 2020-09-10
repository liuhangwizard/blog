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

type VBControllerTag struct {
	VBControllerBase
	modelTag *model.VBModelTag
}


func (c *VBControllerTag) GetTagList(ctx *gin.Context){

	//get all
	data,err:=c.modelTag.DataGetTagList()
	if err !=nil{
		response.QueryError("[Tag][获取所有标签失败]",err,ctx)
		return;
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
		"list":data,
	}))
}

func (c *VBControllerTag) GetTagBlogList(ctx *gin.Context){


	queryData:=request.VBRTagBlogList{}
	if err:=Param.Bind(ctx,&queryData,"[Tag][获取标签文章参数绑定失败]");err!=nil{
		return;
	}


	result,count,err:=c.modelTag.DataTagBlogList(&queryData)
	if err!=nil{
		ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
			"total":0,
			"list":[]interface{}{},
		}))
		return
	}

	//"total":count,
	//"list":append(append(append(append(result,result...),result...),result...),result...),
	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
		"total":count,
		"list":result,
	}))

}

func (c *VBControllerTag) New (ctx *gin.Context){

	queryData:=request.VBRTagNew{}
	if err:=Param.Bind(ctx,&queryData,"[Tag][新增标签参数绑定失败]");err!=nil{
		return;
	}

	if err:=c.modelTag.New(&queryData);err!=nil{
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{}))

}

func (c *VBControllerTag) Delete (ctx *gin.Context){

	queryData:=request.VBRTagDelete{}
	if err:=Param.Bind(ctx,&queryData,"[Tag][删除标签参数绑定失败]");err!=nil{
		return;
	}

	if err:=c.modelTag.Delete(&queryData);err!=nil{
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{}))
}

func NewVBControllerTag() VBControllerTag {

	controllerTag := VBControllerTag{
		modelTag:model.NewModelTag(),
	}
	return controllerTag
}
