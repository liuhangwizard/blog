package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "vb-server/constant"
	"vb-server/model"
	"vb-server/request"
	"vb-server/response"
	VBLog "vb-server/utils/log"
	Param "vb-server/utils/param"
)

type VBControllerBlogRecord struct {
	VBControllerBase
	modelBlogRecord *model.VBModelBlogRecord
}

func NewVBControllerBlogRecord() VBControllerBlogRecord {

	blogRecordController := VBControllerBlogRecord{
		modelBlogRecord: model.NewModelBlogRecord(),
	}
	return blogRecordController
}

func (c *VBControllerBlogRecord) RecordNew(ctx *gin.Context){

	queryData := request.VBRBlogRecordNew{}
	if err:=Param.Bind(ctx,&queryData,"[BlogRecord][新增记录参数绑定失败]");err!=nil{
		return;
	}

	//success
	if err:=c.modelBlogRecord.DataRecordNew(&queryData);err!=nil{
		ctx.JSON(http.StatusOK, response.Success(COMMON_ERROR, gin.H{}))
		return;
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{}))

}

func (c *VBControllerBlogRecord) RecordInfo(ctx *gin.Context) {

	result, err := c.modelBlogRecord.DataRecordInfo()
	if err != nil {
		VBLog.Print.Info("[BlogRecord][平台记录信息获取失败或者为空]" + err.Error())
	}


	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"info": result,
	}))

}

func (c *VBControllerBlogRecord) RecordChartBlog(ctx *gin.Context) {

	//基本信息
	dataBasic, err := c.modelBlogRecord.DataChartBlogBasic()
	if err != nil {
		VBLog.Print.Info("[BlogRecord][Basic信息获取失败]" + err.Error())
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//热度信息
	dataHot, err := c.modelBlogRecord.DataChartBlogHot()
	if err != nil {
		VBLog.Print.Info("[BlogRecord][Hot信息获取失败]" + err.Error())
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	if len(dataHot) != len(dataBasic) {
		VBLog.Print.Info("[BlogRecord][整合BaiscHot信息获取失败]")
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//result
	var result = make([]response.VBResultBlogRecordChartBlog, len(dataHot))
	for i, _ := range dataHot {
		result[i].Id = dataBasic[i].Id
		result[i].Title = dataBasic[i].Title
		result[i].Catagory = dataBasic[i].Catagory
		result[i].Length = dataBasic[i].Length
		result[i].Record = dataBasic[i].Record

		result[i].Hot = dataHot[i].Hot
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"info": result,
	}))

}

func (c *VBControllerBlogRecord) QueryParam(ctx *gin.Context) {

	queryData := request.VBRBlogRecordQueryParam{}
	if err:=Param.Bind(ctx,&queryData,"[BlogRecord][文章记录查询参数绑定失败]");err!=nil{
		return;
	}

	result, count, err := c.modelBlogRecord.DataQueryParam(&queryData)
	if err != nil {
		VBLog.Print.Info("[BlogRecord][文章记录查询失败]" + err.Error())
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"total": count,
		"list":  result,
	}))
}

func (c *VBControllerBlogRecord) ClientMainCount(ctx *gin.Context){

	//count
	resultCount:=c.modelBlogRecord.DataMainCount()

	//propose
	modelBlog:=model.NewModelBlog()
	proposeRequest:=request.VBRBlogInfoPropose{
		CurrentId: 0,
		Catagory:  "",
		Order:     "create_date DESC",
		Max:       5,
	}
	resultPropose,err:=modelBlog.DataClientInfoPropose(&proposeRequest)

	if err != nil {
		VBLog.Print.Info("[BlogRecord][文章CountPropose查询失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"count": resultCount,
		"propose":resultPropose,
	}))
}