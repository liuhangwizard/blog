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
	. "vb-server/utils/session"
)

type VBControllerBlog struct {
	VBControllerBase
	modelBlog      *model.VBModelBlog
	modelBlogImage *model.VBModelBlogImage
}

func NewVBControllerBlog() VBControllerBlog {
	//new controller
	blogController := VBControllerBlog{
		modelBlog:      model.NewModelBlog(),
		modelBlogImage: model.NewModelBlogImage(),
	}
	return blogController
}

func (c *VBControllerBlog) New(ctx *gin.Context) {

	//session user
	userData := ctx.MustGet("User").(VBSessionData)

	artilceData := request.VBRBlogNew{}
	if err := Param.Bind(ctx, &artilceData, "[Blog][文章创建参数绑定失败]"); err != nil {
		return
	}

	err := c.modelBlog.DataNewBlog(&artilceData, userData.UserId)
	if err != nil {
		//error request param
		VBLog.Print.Info("[Blog][文章创建失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(BLOG_SUCCESS_NEW, map[string]interface{}{}))
}

func (c *VBControllerBlog) QueryParam(ctx *gin.Context) {

	//session user
	userData := ctx.MustGet("User").(VBSessionData)

	queryData := request.VBRBlogQueryParam{}
	if err := Param.Bind(ctx, &queryData, "[Blog][文章查询参数绑定失败]"); err != nil {
		return
	}

	result, count, err := c.modelBlog.DataQueryParam(&queryData, userData.UserId)
	if err != nil {
		VBLog.Print.Info("[Blog][文章参数查询失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"total": count,
		"list":  result,
	}))

}

func (c *VBControllerBlog) Update(ctx *gin.Context) {

	queryData := request.VBRBlogUpdate{}
	if err := Param.Bind(ctx, &queryData, "[Blog][文章更新参数绑定失败]"); err != nil {
		return
	}

	if err := c.modelBlog.DataUpdateBlog(&queryData); err != nil {
		VBLog.Print.Info("[Blog][文章更新失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{}))
}

func (c *VBControllerBlog) UpdateStatus(ctx *gin.Context) {

	queryData := request.VBRBlogUpdateStatus{}
	if err := Param.Bind(ctx, &queryData, "[Blog][文章更新状态参数绑定失败]"); err != nil {
		return
	}

	if err := c.modelBlog.DataUpdateStatus(&queryData); err != nil {
		VBLog.Print.Info("[Blog][文章更新状态失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{}))

}

func (c *VBControllerBlog) ManageInfo(ctx *gin.Context) {

	queryData := request.VBRBlogInfo{}
	if err := Param.Bind(ctx, &queryData, "[Blog][文章查询信息参数绑定失败]"); err != nil {
		return
	}

	//查询主体信息
	data, err := c.modelBlog.DataManageInfo(&queryData)
	if err != nil {
		VBLog.Print.Info("[Blog][文章查询信息失败或者为空]" + err.Error())
	}

	//查询图片信息
	imageQueryData:=request.VBRBlogInfoImage{
		BlogId:data.Id,
	}
	imageData,err:=c.modelBlogImage.DataBlogImageInfo(&imageQueryData)
	if err!=nil{
		VBLog.Print.Info("[Blog][文章图片信息查询失败或者为空]" + err.Error())
	}

	//填充image信息
	data.ImageId=imageData.ImageId

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"info": data,
	}))

}

func (c *VBControllerBlog) ClientArticleInfo(ctx *gin.Context) {

	queryData := request.VBRBlogInfo{}
	if err := Param.Bind(ctx, &queryData, "[Blog][Client文章查询信息参数绑定失败]"); err != nil {
		return
	}

	data, err := c.modelBlog.DataClientInfo(&queryData)
	if err != nil {
		VBLog.Print.Info("[Blog][Client文章查询信息失败或者为空]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//propose
	articlePropose := request.VBRBlogInfoPropose{
		CurrentId: data.Id,
		Catagory:  data.Catagory,
		Order:     "create_date DESC",
		Max:       5,
	}

	dataPropose, errPropose := c.modelBlog.DataClientInfoPropose(&articlePropose)
	if errPropose != nil {
		VBLog.Print.Info("[Blog][Client-Propose文章查询信息失败或者为空]" + errPropose.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"info":    data,
		"propose": dataPropose,
	}))

}

//客户端的
func (c *VBControllerBlog) ClientArticlePropose(ctx *gin.Context) {

	queryData := request.VBRBlogInfoPropose{}
	if err := Param.Bind(ctx, &queryData, "[Blog][ClientPropse文章查询信息参数绑定失败]"); err != nil {
		return
	}

	data, err := c.modelBlog.DataClientInfoPropose(&queryData)
	if err != nil {
		VBLog.Print.Info("[Blog][ClientPropse文章查询信息失败或者为空]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"info": data,
	}))

}

func (c *VBControllerBlog) ClientMainCloud(ctx *gin.Context) {

	//archive
	dataArchive, err := c.modelBlog.DataClientMainArchive()
	if err != nil {
		VBLog.Print.Info("[Blog][Cloud归档信息获取失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
	}

	//model
	modelTag := model.NewModelTag()

	//tag
	dataTag, err := modelTag.DataGetTagList()
	if err != nil {
		VBLog.Print.Info("[Blog][Cloud标签信息获取失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
	}

	//catagory
	//modelCatagory:=model.NewModelCatagory()
	//dataCatagory,err:=modelCatagory.DataGetCatagoryList()
	//if err!=nil{
	//	VBLog.Print.Info("[Blog][Cloud分类信息获取失败]" + err.Error())
	//	ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
	//}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"archive": dataArchive,
		"tag":     dataTag,
	}))
}

func (c *VBControllerBlog) ClientMainRecent(ctx *gin.Context) {

	queryData := request.VBRBlogMainRecent{}
	if err := Param.Bind(ctx, &queryData, "[Blog][BlogRecent参数绑定失败]"); err != nil {
		return
	}

	resultBlog, countBlog, err := c.modelBlog.DataClientMainRecent(&queryData)
	if err != nil {
		VBLog.Print.Info("[Blog][BlogRecent查询失败]" + err.Error())
		ctx.JSON(http.StatusInternalServerError, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//整理文章id列表
	blogIdList := make([]uint32, len(resultBlog))
	for i, v := range resultBlog {
		blogIdList[i]=v.Id
	}
	blogImageRequest:=request.VBRBlogMainRecentImageList{
		BlogIdList:blogIdList,
	}

	resultBlogImage,err:=c.modelBlogImage.DataBlogRecentImageList(&blogImageRequest)

	if len(resultBlogImage)==len(resultBlog){
		for i, _ := range resultBlog {
			resultBlog[i].Image=resultBlogImage[i].Image
			//fmt.Println("图片"+strconv.Itoa(int(resultBlog[i].Id)),resultBlog[i].Image)
		}

	}else{
		//如果不相等应该有逻辑错误 这里就不映射图片了
		VBLog.Print.Info("[Blog][BlogRecent文章长度和图片长度不匹配]")

	}




	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"total": countBlog,
		"list":  resultBlog,
	}))

}
