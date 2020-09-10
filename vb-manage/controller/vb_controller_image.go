package controller

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	. "vb-server/constant"
	VBOSS "vb-server/database/oss"
	"vb-server/model"
	"vb-server/request"
	"vb-server/response"
	"vb-server/utils/config"
	VBLog "vb-server/utils/log"
	Param "vb-server/utils/param"
)

type VBControllerImage struct {
	VBControllerBase
	modelImage *model.VBModelImage
	ossInfo *config.VBConfigDBOSS
}

func NewVBControllerImage() VBControllerImage {
	//new controller
	info:=config.VBGetConfigDBOSS();
	controllerImage := VBControllerImage{
		modelImage:model.NewModelImage(),
		ossInfo:&info,
	}
	return controllerImage
}

func (c *VBControllerImage) New(ctx *gin.Context){

	//oss上传成功后，前端返回的文件信息
	queryData:=request.VBRImageNew{}
	if err:=Param.Bind(ctx,&queryData,"[Image][新增图片参数绑定失败]");err!=nil{
		return;
	}

	//图片
	err:=c.modelImage.DataNewImage(&queryData)
	if err!=nil{
		VBLog.Print.Info("[Image][新增图片失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(IMAGE_ERROR_FILEBACK, map[string]interface{}{}))
	}

	ctx.JSON(http.StatusOK, response.Success(IMAGE_SUCCESS_NEW, map[string]interface{}{}))

}

func (c *VBControllerImage) UploadSign(ctx *gin.Context){


	//oss配置文件,获取默认的上传路径
	queryData:=request.VBRImageNewSign{}
	if err:=Param.Bind(ctx,&queryData,"[Image][获取签名参数绑定失败]");err!=nil{
		return;
	}

	//上传的文件扩展名 用于设置Content-Type
	//必须设定,这个用于计算签名
	ext := filepath.Ext(queryData.FileName)

	//设置Content-Type
	typeName:=mime.TypeByExtension(ext);
	options := []oss.Option{
		oss.ContentType(typeName),
	}

	//options := []oss.Option{
	//	oss.ContentType("multipart/form-data"),
	//}

	objectKey := c.ossInfo.Dirctory + "/" + uuid.NewV4().String() + ext


	signedPutURL, err :=VBOSS.Bucket.SignURL(objectKey,oss.HTTPPut,c.ossInfo.Expire,options...)
	if err!=nil{
		VBLog.Print.Info("[Image][签名失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(IMAGE_ERROR_SIGN_FAIL, map[string]interface{}{}))
	}


	//如果启用自定义域名 需要处理一下
	finalURL:=strings.Replace(signedPutURL,c.ossInfo.Bucket+".","",-1)

	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS, gin.H{
		"upload":finalURL,
		"type":typeName,
		"address":objectKey,
	}))

}

func (c *VBControllerImage) QueryParam(ctx *gin.Context){

	queryData:=request.VBRImageQueryParam{}
	if err:=Param.Bind(ctx,&queryData,"[Image][查询图片参数绑定失败]");err!=nil{
		return;
	}


	result,count,err:=c.modelImage.DataQueryParam(&queryData)
	if err!=nil{
		VBLog.Print.Info("[Blog][文章参数查询失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(COMMON_ERROR, map[string]interface{}{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{
		"total":count,
		"list":result,
	}))
}

func (c *VBControllerImage) UpdateType(ctx *gin.Context){

	queryData:=request.VBRImageUpdateType{}
	if err:=Param.Bind(ctx,&queryData,"[Image][修改类型参数绑定失败]");err!=nil{
		return;
	}

	err:=c.modelImage.DataUpateType(&queryData)
	if err !=nil{
		VBLog.Print.Info("[Image][更新参数失败]" + err.Error())
		ctx.JSON(http.StatusOK, response.Error(IMAGE_ERROR_TYPE_FAIL, gin.H{}))
		return
	}

	//success
	ctx.JSON(http.StatusOK, response.Success(COMMON_SUCCESS,gin.H{}))
}